package product

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductController interface {
	CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
