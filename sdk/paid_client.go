package sdk

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/AgentPaid/ap-go-client/models"
)

var (
	DefaultAPIURL = "https://api.agentpaid.io"
)

type PaidClient struct {
	apiKey    string
	apiURL    string
	signals   []any
	mu        sync.Mutex
	client    *http.Client
	ticker    *time.Ticker
	quit      chan struct{}
	flushFreq time.Duration
	logger    *slog.Logger
}

func NewPaidClient(apiKey string, apiURL ...string) *PaidClient {
	if apiKey == "" {
		panic("api key is required")
	}

	url := DefaultAPIURL
	if len(apiURL) > 0 && apiURL[0] != "" {
		url = apiURL[0]
	}

	c := &PaidClient{
		apiKey:    apiKey,
		apiURL:    url,
		signals:   make([]any, 0),
		client:    &http.Client{Timeout: 10 * time.Second},
		flushFreq: 30 * time.Second, // Hardcoded 30s interval
		quit:      make(chan struct{}),
		logger:    slog.Default(),
	}
	c.startTimer()
	return c
}

func (c *PaidClient) startTimer() {
	c.ticker = time.NewTicker(c.flushFreq)
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.Flush()
			case <-c.quit:
				return
			}
		}
	}()
}

func (c *PaidClient) RecordUsage(agentID, customerID, eventName string, data any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	signal := models.Signal[any]{
		EventName:  eventName,
		AgentID:    agentID,
		CustomerID: customerID,
		Data:       data,
	}

	c.signals = append(c.signals, signal)

	if len(c.signals) >= 100 {
		log.Println("Buffer reached 100, auto-flushing...")
		c.Flush()
	}
}

func (c *PaidClient) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.signals) == 0 {
		return
	}

	body := map[string]any{"transactions": c.signals}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.logger.Error("Error marshalling signals: %v", "error", err)
		return
	}

	req, err := http.NewRequest("POST", c.apiURL+"/api/entries/bulk", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.logger.Error("Request creation failed: %v", "error", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("Flush failed: %v", "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		c.logger.Info("Flushed signals successfully", "count", len(c.signals))
		c.signals = make([]any, 0)
	} else {
		log.Printf("Flush failed with status %d", resp.StatusCode)
	}
}

func (c *PaidClient) Close() {
	close(c.quit)
	c.ticker.Stop()
	c.logger.Info("Closing client, flushing remaining signals...")
	c.Flush()
}
