package order_product

import (
	"context"
	order_product2 "go-rest-api/model/web/order-product"
)

type OrderProductServiceRe interface {
	CreateOrderProduct(ctx context.Context, request order_product2.OrderProductCreateRequest) order_product2.OrderProductResponse
	UpdateOrderProduct(ctx context.Context, request order_product2.OrderProductUpdateRequest) order_product2.OrderProductResponse
	DeleteOrderProduct(ctx context.Context, orderproductId int)
	FindOrderProductById(ctx context.Context, orderproductId int) order_product2.OrderProductResponse
	FindAllOrderProduct(ctx context.Context) []order_product2.OrderProductResponse
}
