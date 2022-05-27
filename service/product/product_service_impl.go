package product

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/helper/model"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/product"
	product2 "go-rest-api/repository/product"
)

type ProductServiceImpl struct {
	ProductRepository product2.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository product2.ProductRepository, DB *sql.DB, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, request product.ProductCreateRequest) product.ProductResponse {
	logrus.Info("Product Service Create Start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}

	product = service.ProductRepository.SaveProduct(ctx, tx, product)

	logrus.Info("Product Service Create End")
	return model.ToProductResponse(product)
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, request product.ProductUpdateRequest) product.ProductResponse {
	logrus.Info("Product Service Update Start")
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	productResponse, err := service.ProductRepository.FindProductById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	products := model.ToProduct(productResponse)
	products.Name = request.Name
	products.Price = request.Price
	products.CategoryId = request.CategoryId

	products = service.ProductRepository.UpdateProduct(ctx, tx, products)

	logrus.Info("Product Service Update End")
	return model.ToProductResponse(products)
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, productId int) {
	logrus.Info("Product Service Delete Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	productResponse, err := service.ProductRepository.FindProductById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := model.ToProduct(productResponse)

	service.ProductRepository.DeleteProduct(ctx, tx, product)
	logrus.Info("Product Service Delete End")
}

func (service *ProductServiceImpl) FindProductById(ctx context.Context, productId int) product.ProductResponse {
	logrus.Info("Product Service Delete Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	product, err := service.ProductRepository.FindProductById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	logrus.Info("Product Service Delete End")
	/*	return model.ToProductResponse(product)*/
	return product
}

func (service *ProductServiceImpl) FindAllProduct(ctx context.Context) []product.ProductResponse {
	logrus.Info("Product Service Find All Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	product := service.ProductRepository.FindAllProduct(ctx, tx)

	logrus.Info("Product Service Find All End")
	/*	return model.ToProductResponses(product)*/
	return product
}
