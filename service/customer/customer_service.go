package customer

import (
	"context"
	"go-rest-api/model/web/customer"
)

type CustomerService interface {
	CreateCustomer(ctx context.Context, request customer.CustomerCreateRequest) customer.CustomerResponse
	UpdateCustomer(ctx context.Context, request customer.CustomerUpdateRequest) customer.CustomerResponse
	DeleteCustomer(ctx context.Context, customersId int)
	FindCustomerById(ctx context.Context, customersId int) customer.CustomerResponse
	FindAllCustomer(ctx context.Context) []customer.CustomerResponse
}
