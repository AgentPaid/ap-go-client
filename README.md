# AgentPaid Go Client

A Go client for interacting with the Paid API. This package allows you to send transaction events to the API.

## Installation

Use the Go package manager to install the client:

```bash
go get github.com/AgentPaid/ap-go-client
```

## Usage

```go
package main

import (
    "github.com/AgentPaid/ap-go-client/sdk"
)

func main() {
    // Initialize the client
    client := sdk.NewPaidClient("your-api-key")
    defer client.Close() // Ensure proper cleanup

    // Record usage events
    client.RecordUsage(
        "your-agent-id",
        "customer-id",
        "event-name",
        map[string]any{
            "key": "data",
        },
    )

    // Events are automatically flushed:
    // - Every 30 seconds
    // - When the buffer reaches 100 events
    // To manually flush:
    client.Flush()
}
```

## Features

- Automatic event batching
- Periodic flushing (every 30 seconds)
- Buffer-based flushing (at 100 events)
- Thread-safe operations
- Graceful shutdown handling
