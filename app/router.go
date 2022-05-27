package app

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/controller/category"
	"go-rest-api/controller/customer"
	order_product "go-rest-api/controller/order-product"
	"go-rest-api/controller/orders"
	"go-rest-api/controller/product"
	"go-rest-api/exception"
)

func NewRouter(categoryController category.CategoryController, ordersController orders.OrdersController, customerController customer.CustomerController, productController product.ProductController, orderProductController order_product.OrderProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAllCategory)
	router.GET("/api/categories/:categoryId", categoryController.FindCategoryById)
	router.POST("/api/categories", categoryController.CreateCategory)
	router.PUT("/api/categories/:categoryId", categoryController.UpdateCategory)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteCategory)

	router.GET("/api/orderss", ordersController.FindAllOrders)
	router.GET("/api/orderss/:ordersId", ordersController.FindOrdersById)
	router.POST("/api/orderss", ordersController.CreateOrders)
	router.PUT("/api/orderss/:ordersId", ordersController.UpdateOrders)
	router.DELETE("/api/orderss/:ordersId", ordersController.DeleteOrders)

	router.GET("/api/customers", customerController.FindAllCustomer)
	router.GET("/api/customers/:customerId", customerController.FindCustomerById)
	router.POST("/api/customers", customerController.CreateCustomer)
	router.PUT("/api/customers/:customerId", customerController.UpdateCustomer)
	router.DELETE("/api/customers/:customerId", customerController.DeleteCustomer)

	router.GET("/api/products", productController.FindAllProduct)
	router.GET("/api/products/:productId", productController.FindProductById)
	router.POST("/api/products", productController.CreateProduct)
	router.PUT("/api/products/:productId", productController.UpdateProduct)
	router.DELETE("/api/products/:productId", productController.DeleteProduct)

	router.GET("/api/orderProducts", orderProductController.FindAllOrderProduct)
	router.GET("/api/orderProducts/:orderProductId", orderProductController.FindOrderProductById)
	router.POST("/api/orderProducts", orderProductController.CreateOrderProduct)
	router.PUT("/api/orderProducts/:orderProductId", orderProductController.UpdateOrderProduct)
	router.DELETE("/api/orderProducts/:orderProductId", orderProductController.DeleteOrderProduct)

	router.PanicHandler = exception.ErrorHandler

	return router
}
