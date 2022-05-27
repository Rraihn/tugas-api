package category

import (
	"context"
	"go-rest-api/model/web/category"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, request category.CategoryCreateRequest) category.CategoryResponse
	UpdateCategory(ctx context.Context, request category.CategoryUpdateRequest) category.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
	FindCategoryById(ctx context.Context, categoryId int) category.CategoryResponse
	FindAllCategory(ctx context.Context) []category.CategoryResponse
}
