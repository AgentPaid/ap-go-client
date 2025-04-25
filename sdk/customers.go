package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AgentPaid/ap-go-client/models"
)

func (c *PaidClient) CreateCustomer(orgID string, customer models.CreateCustomerRequest) (*models.Customer, error) {
	c.logger.Info("Creating customer", "orgID", orgID)

	jsonBody, err := json.Marshal(customer)
	if err != nil {
		c.logger.Error("Error marshalling customer", "error", err)
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/organizations/%s/customers", c.apiURL, orgID),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		c.logger.Error("Request creation failed", "error", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("Create customer request failed", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response struct {
			Data    models.Customer `json:"data"`
			Message string          `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}
		return &response.Data, nil
	}

	return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
}

func (c *PaidClient) ListCustomers(orgID string) (*models.ListCustomersResponse, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/api/organizations/%s/customers", c.apiURL, orgID),
		nil,
	)
	if err != nil {
		c.logger.Error("Request creation failed", "error", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("List customers request failed", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response struct {
			Data    []models.Customer `json:"data"`
			Message string            `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		return &models.ListCustomersResponse{
			Customers: response.Data,
		}, nil
	}

	var errorResponse struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, errorResponse.Message)
}

func (c *PaidClient) GetCustomer(orgID, customerID string) (*models.Customer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/api/organizations/%s/customer/%s", c.apiURL, orgID, customerID),
		nil,
	)
	if err != nil {
		c.logger.Error("Request creation failed", "error", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("Get customer request failed", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response struct {
			Data    models.Customer `json:"data"`
			Message string          `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		return &response.Data, nil
	}

	var errorResponse struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, errorResponse.Message)
}

func (c *PaidClient) UpdateCustomer(orgID, customerID string, request models.UpdateCustomerRequest) (*models.Customer, error) {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		c.logger.Error("Error marshalling customer update", "error", err)
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/api/organizations/%s/customer/%s", c.apiURL, orgID, customerID),
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		c.logger.Error("Request creation failed", "error", err)
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("Update customer request failed", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response struct {
			Data    models.Customer `json:"data"`
			Message string          `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			c.logger.Error("Failed to decode response", "error", err)
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		return &response.Data, nil
	}

	var errorResponse struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, errorResponse.Message)
}

func (c *PaidClient) DeleteCustomer(orgID, customerID string) error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/api/organizations/%s/customer/%s", c.apiURL, orgID, customerID),
		nil,
	)
	if err != nil {
		c.logger.Error("Request creation failed", "error", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("Delete customer request failed", "error", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	return fmt.Errorf("request failed with status %d", resp.StatusCode)
}
