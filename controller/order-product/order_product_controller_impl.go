package order_product

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/model/web/order-product"
	order_product2 "go-rest-api/service/order-product"
	"net/http"
	"strconv"
)

type OrderProductControllerImpl struct {
	OrderProductService order_product2.OrderProductServiceRe
}

func NewOrderProductController(orderProductService order_product2.OrderProductServiceRe) OrderProductController {
	return &OrderProductControllerImpl{
		OrderProductService: orderProductService,
	}
}

func (o OrderProductControllerImpl) CreateOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductCreateRequest := order_product.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &orderProductCreateRequest)

	orderProductResponse := o.OrderProductService.CreateOrderProduct(request.Context(), orderProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (o OrderProductControllerImpl) UpdateOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductUpdateRequest := order_product.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderProductUpdateRequest)

	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	orderProductUpdateRequest.Id = id

	orderProductResponse := o.OrderProductService.UpdateOrderProduct(request.Context(), orderProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   orderProductResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (o OrderProductControllerImpl) DeleteOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	o.OrderProductService.DeleteOrderProduct(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (o OrderProductControllerImpl) FindOrderProductById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductId := params.ByName("orderProductId")
	id, err := strconv.Atoi(orderProductId)
	helper.PanicIfError(err)

	orderProductResponse := o.OrderProductService.FindOrderProductById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (o OrderProductControllerImpl) FindAllOrderProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductResponse := o.OrderProductService.FindAllOrderProduct(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
