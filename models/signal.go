package models

type Signal[T any] struct {
	EventName  string `json:"event_name"`
	AgentID    string `json:"agent_id"`
	CustomerID string `json:"customer_id"`
	Data       T      `json:"data"`
}
