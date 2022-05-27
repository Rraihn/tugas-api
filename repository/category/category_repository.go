package category

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type CategoryRepository interface {
	SaveCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	UpdateCategory(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	DeleteCategory(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindCategoryById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAllCategory(ctx context.Context, tx *sql.Tx) []domain.Category
}
