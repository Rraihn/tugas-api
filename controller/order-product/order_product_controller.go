package order_product

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OrderProductController interface {
	CreateOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
