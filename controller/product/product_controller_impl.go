package product

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/model/web/product"
	product2 "go-rest-api/service/product"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService product2.ProductServiceRe
}

func NewProductController(ProductService product2.ProductServiceRe) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (p ProductControllerImpl) CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Create Product Start")
	productCreateRequest := product.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := p.ProductService.CreateProduct(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("Create Product End")
}

func (p ProductControllerImpl) UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Update Product Start")
	productUpdateRequest := product.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := p.ProductService.UpdateProduct(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("Update Product End")
}

func (p ProductControllerImpl) DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Delete Product Start")
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	p.ProductService.DeleteProduct(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("Delete Product End")
}

func (p ProductControllerImpl) FindProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Find By Id Start")
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := p.ProductService.FindProductById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("Find By Id End")
}

func (p ProductControllerImpl) FindAllProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	logrus.Info("Find All Start")
	productResponse := p.ProductService.FindAllProduct(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
	logrus.Info("Find All End")
}
