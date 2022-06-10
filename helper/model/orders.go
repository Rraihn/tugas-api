package model

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web/orders"
	"time"
)

func ToOrdersResponse(ord domain.Orders) orders.OrdersResponse {
	return orders.OrdersResponse{
		Id:         ord.Id,
		CustomerId: ord.CustomerID,
		DateTime:   time.Now(),
	}
}

func ToOrdersResponses(ord []domain.Orders) []orders.OrdersResponse {
	var ordersResponses []orders.OrdersResponse
	for _, ordeer := range ord {
		ordersResponses = append(ordersResponses, ToOrdersResponse(ordeer))
	}
	return ordersResponses
}
