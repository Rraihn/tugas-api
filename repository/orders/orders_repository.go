package orders

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type OrdersRepository interface {
	SaveOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	UpdateOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders
	DeleteOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders)
	FindOrdersById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error)
	FindAllOrders(ctx context.Context, tx *sql.Tx) []domain.Orders
}
