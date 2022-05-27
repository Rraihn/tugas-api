package customer

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (c CustomerRepositoryImpl) SaveCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "insert into customer(name, address, email, phone_number, createdAt) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Email, customer.Address, customer.PhoneNumber, customer.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer
}

func (c CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) domain.Customer {
	SQL := "update customer set name = ?, email = ?, address = ?, phone_number = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Name, customer.Email, customer.Address, customer.PhoneNumber, customer.UpdatedAt, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (c CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, tx *sql.Tx, customer domain.Customer) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (c CustomerRepositoryImpl) FindCustomerById(ctx context.Context, tx *sql.Tx, customerId int) (domain.Customer, error) {
	SQL := "select id, name, address, email, phone_number from customer where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}

	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		defer rows.Close()
		return customer, nil
	} else {
		return customer, errors.New("Customer is not found")
	}
}

func (c CustomerRepositoryImpl) FindAllCustomer(ctx context.Context, tx *sql.Tx) []domain.Customer {
	SQL := "select id, name, address, email, phone_number from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Address, &customer.Email, &customer.PhoneNumber)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}
	return customers
}
