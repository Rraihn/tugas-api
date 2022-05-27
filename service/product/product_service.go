package product

import (
	"context"
	"go-rest-api/model/web/product"
)

type ProductServiceRe interface {
	CreateProduct(ctx context.Context, request product.ProductCreateRequest) product.ProductResponse
	UpdateProduct(ctx context.Context, request product.ProductUpdateRequest) product.ProductResponse
	DeleteProduct(ctx context.Context, productId int)
	FindProductById(ctx context.Context, productId int) product.ProductResponse
	FindAllProduct(ctx context.Context) []product.ProductResponse
}
