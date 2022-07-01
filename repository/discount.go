package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type DiscountRepository interface {
	SaveDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	UpdateDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount
	DeleteDiscount(ctx context.Context, tx *sql.Tx, discount domain.Discount)
	FindDiscountById(ctx context.Context, tx *sql.Tx, discountId int)
	FindAllDiscount(ctx context.Context, tx *sql.Tx) []discount
}
