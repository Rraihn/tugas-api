package order_product

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/helper/model"
	"go-rest-api/model/domain"
	order_product "go-rest-api/model/web/order-product"
	order_product2 "go-rest-api/repository/order-product"
)

type OrderProductServiceImpl struct {
	OrderProductRepository order_product2.OrderProductRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewOrderProductService(orderProductRepository order_product2.OrderProductRepository, DB *sql.DB, validate *validator.Validate) *OrderProductServiceImpl {
	return &OrderProductServiceImpl{
		OrderProductRepository: orderProductRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service OrderProductServiceImpl) CreateOrderProduct(ctx context.Context, request order_product.OrderProductCreateRequest) order_product.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderProduct := domain.OrderProduct{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Price:     request.Price,
		Qty:       request.Qty,
	}

	orderProduct = service.OrderProductRepository.SaveOrderProduct(ctx, tx, orderProduct)

	return model.ToOrderProductResponse(orderProduct)
}

func (service OrderProductServiceImpl) UpdateOrderProduct(ctx context.Context, request order_product.OrderProductUpdateRequest) order_product.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderProduct, err := service.OrderProductRepository.FindOrderProductById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	orderProduct.Id = request.Id
	orderProduct.OrderId = request.OrderId
	orderProduct.ProductId = request.ProductId
	orderProduct.Price = request.Price
	orderProduct.Qty = request.Qty

	orderProduct = service.OrderProductRepository.UpdateOrderProduct(ctx, tx, orderProduct)

	return model.ToOrderProductResponse(orderProduct)
}

func (service OrderProductServiceImpl) DeleteOrderProduct(ctx context.Context, orderproductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderProduct, err := service.OrderProductRepository.FindOrderProductById(ctx, tx, orderproductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderProductRepository.DeleteOrderProduct(ctx, tx, orderProduct)
}

func (service OrderProductServiceImpl) FindOrderProductById(ctx context.Context, orderproductId int) order_product.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderProduct, err := service.OrderProductRepository.FindOrderProductById(ctx, tx, orderproductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToOrderProductResponse(orderProduct)
}

func (service OrderProductServiceImpl) FindAllOrderProduct(ctx context.Context) []order_product.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	orderProduct := service.OrderProductRepository.FindAllOrderProduct(ctx, tx)

	return model.ToOrderProductResponses(orderProduct)
}
