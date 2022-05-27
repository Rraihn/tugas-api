package order_product

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type OrderProductRepository interface {
	SaveOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct
	UpdateOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct
	DeleteOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct)
	FindOrderProductById(ctx context.Context, tx *sql.Tx, orderproductId int) (domain.OrderProduct, error)
	FindAllOrderProduct(ctx context.Context, tx *sql.Tx) []domain.OrderProduct
}
