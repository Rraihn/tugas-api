package orders

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OrdersController interface {
	CreateOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrdersById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
