package orders

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/model/web/orders"
	orders2 "go-rest-api/service/orders"
	"net/http"
	"strconv"
)

type OrdersControllerImpl struct {
	OrdersService orders2.OrdersService
}

func NewOrdersController(ordersService orders2.OrdersService) OrdersController {
	return &OrdersControllerImpl{
		OrdersService: ordersService,
	}
}

func (ord OrdersControllerImpl) CreateOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersCreateRequest := orders.OrdersCreateRequest{}
	helper.ReadFromRequestBody(request, &ordersCreateRequest)

	ordersResponse := ord.OrdersService.CreateOrders(request.Context(), ordersCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ord OrdersControllerImpl) UpdateOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersUpdateRequest := orders.OrdersUpdateRequest{}
	helper.ReadFromRequestBody(request, &ordersUpdateRequest)

	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	ordersUpdateRequest.Id = id

	ordersResponse := ord.OrdersService.UpdateOrders(request.Context(), ordersUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ord OrdersControllerImpl) DeleteOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	ord.OrdersService.DeleteOrders(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ord OrdersControllerImpl) FindOrdersById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersId := params.ByName("ordersId")
	id, err := strconv.Atoi(ordersId)
	helper.PanicIfError(err)

	ordersResponse := ord.OrdersService.FindOrdersById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   ordersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (ord OrdersControllerImpl) FindAllOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ordersResponse := ord.OrdersService.FindAllOrders(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
