package product

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/product"
)

type ProductRepository interface {
	SaveProduct(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	UpdateProduct(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	DeleteProduct(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindProductById(ctx context.Context, tx *sql.Tx, productId int) (product.ProductResponse, error)
	FindAllProduct(ctx context.Context, tx *sql.Tx) []product.ProductResponse
}
