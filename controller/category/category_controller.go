package category

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	CreateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
