package customer

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type CustomerRepository interface {
	SaveCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	UpdateCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer
	DeleteCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer)
	FindCustomerById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error)
	FindAllCustomer(ctx context.Context, tx *sql.Tx) []domain.Customer
}
