package sdk

type Client interface {
	RecordUsage(agentID, customerID, eventName string, data any)
	Flush()
	Close()
}
