package sdk

import (
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var (
	DefaultAPIURL = "https://api.agentpaid.io"
)

type PaidClient struct {
	apiKey    string
	apiURL    string
	signals   []any
	mu        sync.RWMutex
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

func (c *PaidClient) Close() {
	close(c.quit)
	c.ticker.Stop()
	c.logger.Info("Closing client, flushing remaining signals...")
	c.Flush()
}
