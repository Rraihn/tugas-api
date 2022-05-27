package orders

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/helper/model"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/orders"
	orders2 "go-rest-api/repository/orders"
)

type OrdersServiceImpl struct {
	OrdersRepository orders2.OrdersRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewOrdersService(ordersRepository orders2.OrdersRepository, DB *sql.DB, validate *validator.Validate) *OrdersServiceImpl {
	return &OrdersServiceImpl{
		OrdersRepository: ordersRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *OrdersServiceImpl) CreateOrders(ctx context.Context, request orders.OrdersCreateRequest) orders.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderss := domain.Orders{
		CustomerID: request.CustomerId,
	}

	orderss = service.OrdersRepository.SaveOrders(ctx, tx, orderss)

	return model.ToOrdersResponse(orderss)
}

func (service *OrdersServiceImpl) UpdateOrders(ctx context.Context, request orders.OrdersUpdateRequest) orders.OrdersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderss, err := service.OrdersRepository.FindOrdersById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	orderss.CustomerID = request.CustomerId

	orderss = service.OrdersRepository.UpdateOrders(ctx, tx, orderss)

	return model.ToOrdersResponse(orderss)
}

func (service *OrdersServiceImpl) DeleteOrders(ctx context.Context, ordersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderss, err := service.OrdersRepository.FindOrdersById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrdersRepository.DeleteOrders(ctx, tx, orderss)
}

func (service *OrdersServiceImpl) FindOrdersById(ctx context.Context, ordersId int) orders.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderss, err := service.OrdersRepository.FindOrdersById(ctx, tx, ordersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToOrdersResponse(orderss)
}

func (service *OrdersServiceImpl) FindAllOrders(ctx context.Context) []orders.OrdersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderss := service.OrdersRepository.FindAllOrders(ctx, tx)

	return model.ToOrdersResponses(orderss)
}
