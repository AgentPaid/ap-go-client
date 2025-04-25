package models

type Customer struct {
	ID              string   `json:"id"`
	OrganizationID  string   `json:"organizationId"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone,omitempty"`
	EmployeeCount   int      `json:"employeeCount,omitempty"`
	AnnualRevenue   float64  `json:"annualRevenue,omitempty"`
	TaxExemptStatus string   `json:"taxExemptStatus,omitempty"`
	CreationSource  string   `json:"creationSource,omitempty"`
	Website         *string  `json:"website,omitempty"`
	ExternalID      *string  `json:"externalId,omitempty"`
	BillingAddress  *Address `json:"billingAddress,omitempty"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type CreateCustomerRequest struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone,omitempty"`
	EmployeeCount   int      `json:"employeeCount,omitempty"`
	AnnualRevenue   float64  `json:"annualRevenue,omitempty"`
	TaxExemptStatus string   `json:"taxExemptStatus,omitempty"`
	CreationSource  string   `json:"creationSource,omitempty"`
	Website         *string  `json:"website,omitempty"`
	ExternalID      *string  `json:"externalId,omitempty"`
	BillingAddress  *Address `json:"billingAddress,omitempty"`
}

type UpdateCustomerRequest struct {
	Name            string   `json:"name,omitempty"`
	Email           string   `json:"email,omitempty"`
	Phone           string   `json:"phone,omitempty"`
	EmployeeCount   *int     `json:"employeeCount,omitempty"`
	AnnualRevenue   *float64 `json:"annualRevenue,omitempty"`
	TaxExemptStatus *string  `json:"taxExemptStatus,omitempty"`
	CreationSource  string   `json:"creationSource,omitempty"`
	Website         *string  `json:"website,omitempty"`
	ExternalID      *string  `json:"externalId,omitempty"`
	BillingAddress  *Address `json:"billingAddress,omitempty"`
}

type CreateCustomersResponse struct {
	Customers []Customer `json:"customers"`
}

type ListCustomersResponse struct {
	Customers []Customer `json:"customers"`
}
