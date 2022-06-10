package model

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web/category"
	"time"
)

func ToCategoryResponse(c domain.Category) category.CategoryResponse {
	return category.CategoryResponse{
		Id:       c.Id,
		Name:     c.Name,
		DateTime: time.Now(),
	}
}

func ToCategoryResponses(categories []domain.Category) []category.CategoryResponse {
	var categoryResponses []category.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
