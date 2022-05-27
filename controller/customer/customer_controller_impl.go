package customer

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/model/web/customer"
	category2 "go-rest-api/service/customer"
	"net/http"
	"strconv"
)

type CustomerControllerImpl struct {
	CustomerService category2.CustomerService
}

func NewCustomerController(customerService category2.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (c CustomerControllerImpl) CreateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerCreateRequest := customer.CustomerCreateRequest{}
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerResponse := c.CustomerService.CreateCustomer(request.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c CustomerControllerImpl) UpdateCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerUpdateRequest := customer.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(request, &customerUpdateRequest)

	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := c.CustomerService.UpdateCustomer(request.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customerResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (c CustomerControllerImpl) DeleteCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	c.CustomerService.DeleteCustomer(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c CustomerControllerImpl) FindCustomerById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := c.CustomerService.FindCustomerById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c CustomerControllerImpl) FindAllCustomer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerResponse := c.CustomerService.FindAllCustomer(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
