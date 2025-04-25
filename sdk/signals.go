package sdk

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AgentPaid/ap-go-client/models"
)

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
		c.mu.Unlock() // Temporarily unlock to allow Flush to proceed
		c.Flush()
		c.mu.Lock() // Re-lock for the defer
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
