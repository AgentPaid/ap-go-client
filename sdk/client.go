package sdk

import "github.com/AgentPaid/ap-go-client/models"

type Client interface {
	// Signals methods
	RecordUsage(agentID, customerID, eventName string, data any)
	Flush()
	Close()

	// Customer methods
	CreateCustomers(orgID string, customer models.CreateCustomerRequest) (*models.CreateCustomersResponse, error)
	ListCustomers(orgID string) (*models.ListCustomersResponse, error)
	GetCustomer(orgID, customerID string) (*models.Customer, error)
	UpdateCustomer(orgID, customerID string, request models.UpdateCustomerRequest) (*models.Customer, error)
	DeleteCustomer(orgID, customerID string) error
}
