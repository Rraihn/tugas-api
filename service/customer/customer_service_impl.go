package customer

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/helper/model"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/customer"
	category2 "go-rest-api/repository/customer"
	"time"
)

type CustomerServiceImpl struct {
	CustomerRepository category2.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerReposity category2.CustomerRepository, DB *sql.DB, validate *validator.Validate) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		CustomerRepository: customerReposity,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) CreateCustomer(ctx context.Context, request customer.CustomerCreateRequest) customer.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	customer := domain.Customer{
		Name:        request.Name,
		Email:       request.Email,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	customer = service.CustomerRepository.SaveCustomer(ctx, tx, customer)

	return model.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) UpdateCustomer(ctx context.Context, request customer.CustomerUpdateRequest) customer.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)
	
	customer, err := service.CustomerRepository.FindCustomerById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	customer.Name = request.Name
	customer.Address = request.Address
	customer.Email = request.Email
	customer.PhoneNumber = request.PhoneNumber
	customer.UpdatedAt = time.Now()

	customer = service.CustomerRepository.UpdateCustomer(ctx, tx, customer)

	return model.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) DeleteCustomer(ctx context.Context, customersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	customer, err := service.CustomerRepository.FindCustomerById(ctx, tx, customersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.DeleteCustomer(ctx, tx, customer)
}
func (service *CustomerServiceImpl) FindCustomerById(ctx context.Context, customersId int) customer.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	customer, err := service.CustomerRepository.FindCustomerById(ctx, tx, customersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) FindAllCustomer(ctx context.Context) []customer.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	customers := service.CustomerRepository.FindAllCustomer(ctx, tx)

	return model.ToCustomerResponses(customers)
}
