package customer

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CustomerController interface {
	CreateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCustomerById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
