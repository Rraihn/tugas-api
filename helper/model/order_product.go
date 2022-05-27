package model

import (
	"go-rest-api/model/domain"
	order_product "go-rest-api/model/web/order-product"
)

func ToOrderProductResponse(op domain.OrderProduct) order_product.OrderProductResponse {
	return order_product.OrderProductResponse{
		Id:        op.Id,
		OrderId:   op.OrderId,
		ProductId: op.ProductId,
		Qty:       op.Qty,
		Price:     op.Price,
	}
}

func ToOrderProduct(op domain.OrderProduct) order_product.OrderProductResponse {
	return order_product.OrderProductResponse{
		Id:        op.Id,
		OrderId:   op.OrderId,
		ProductId: op.ProductId,
		Qty:       op.Qty,
		Price:     op.Price,
	}
}

func ToOrderProductResponses(op []domain.OrderProduct) []order_product.OrderProductResponse {
	var orderProductResponses []order_product.OrderProductResponse
	for _, orderProduct := range op {
		orderProductResponses = append(orderProductResponses, ToOrderProductResponse(orderProduct))
	}
	return orderProductResponses
}
