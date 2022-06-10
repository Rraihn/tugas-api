package model

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web/customer"
	"time"
)

func ToCustomerResponse(c domain.Customer) customer.CustomerResponse {
	return customer.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		Email:       c.Email,
		Address:     c.Address,
		PhoneNumber: c.PhoneNumber,
		DateTime:    time.Now(),
	}
}

func ToCustomerResponses(customers []domain.Customer) []customer.CustomerResponse {
	var customerResponses []customer.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}
