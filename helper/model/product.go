package model

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web/product"
	"time"
)

func ToProductResponse(p domain.Product) product.ProductResponse {
	return product.ProductResponse{
		Id:         p.Id,
		Name:       p.Name,
		Price:      p.Price,
		CategoryId: p.CategoryId,
		DateTime:   time.Now(),
	}
}

func ToProduct(p product.ProductResponse) domain.Product {
	return domain.Product{
		Id:         p.Id,
		Name:       p.Name,
		Price:      p.Price,
		CategoryId: p.CategoryId,
	}
}

func ToProductResponses(p []domain.Product) []product.ProductResponse {
	var productResponses []product.ProductResponse
	for _, product := range p {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
