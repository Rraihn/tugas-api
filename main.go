package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"go-rest-api/app"
	category3 "go-rest-api/controller/category"
	customer3 "go-rest-api/controller/customer"
	order_product3 "go-rest-api/controller/order-product"
	orders3 "go-rest-api/controller/orders"
	product3 "go-rest-api/controller/product"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/repository/category"
	"go-rest-api/repository/customer"
	order_product "go-rest-api/repository/order-product"
	"go-rest-api/repository/orders"
	"go-rest-api/repository/product"
	category2 "go-rest-api/service/category"
	customer2 "go-rest-api/service/customer"
	order_product2 "go-rest-api/service/order-product"
	orders2 "go-rest-api/service/orders"
	product2 "go-rest-api/service/product"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := category.NewCategoryRepository()
	categoryService := category2.NewCategoryService(categoryRepository, db, validate)
	categoryController := category3.NewCategoryController(categoryService)

	ordersRepo := orders.NewOrdersRepository()
	ordersService := orders2.NewOrdersService(ordersRepo, db, validate)
	ordersController := orders3.NewOrdersController(ordersService)

	customerRepo := customer.NewCustomerRepository()
	customerService := customer2.NewCustomerService(customerRepo, db, validate)
	customerController := customer3.NewCustomerController(customerService)

	productRepo := product.NewProductRepository()
	productService := product2.NewProductService(productRepo, db, validate)
	productController := product3.NewProductController(productService)

	orderProductRepo := order_product.NewOrderProductRepository()
	orderproductService := order_product2.NewOrderProductService(orderProductRepo, db, validate)
	orderProductController := order_product3.NewOrderProductController(orderproductService)

	router := app.NewRouter(categoryController, ordersController, customerController, productController, orderProductController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
