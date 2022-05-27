package orders

import (
	"context"
	"go-rest-api/model/web/orders"
)

type OrdersService interface {
	CreateOrders(ctx context.Context, request orders.OrdersCreateRequest) orders.OrdersResponse
	UpdateOrders(ctx context.Context, request orders.OrdersUpdateRequest) orders.OrdersResponse
	DeleteOrders(ctx context.Context, ordersId int)
	FindOrdersById(ctx context.Context, ordersId int) orders.OrdersResponse
	FindAllOrders(ctx context.Context) []orders.OrdersResponse
}
