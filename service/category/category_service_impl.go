package category

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"go-rest-api/exception"
	"go-rest-api/helper"
	"go-rest-api/helper/model"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/category"
	category2 "go-rest-api/repository/category"
)

type CategoryServiceImpl struct {
	CategoryRepository category2.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryReposity category2.CategoryRepository, DB *sql.DB, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		CategoryRepository: categoryReposity,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, request category.CategoryCreateRequest) category.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.SaveCategory(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, request category.CategoryUpdateRequest) category.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	category, err := service.CategoryRepository.FindCategoryById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	category.Name = request.Name

	category = service.CategoryRepository.UpdateCategory(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	category, err := service.CategoryRepository.FindCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.DeleteCategory(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindCategoryById(ctx context.Context, categoryId int) category.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	category, err := service.CategoryRepository.FindCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAllCategory(ctx context.Context) []category.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	categories := service.CategoryRepository.FindAllCategory(ctx, tx)

	return model.ToCategoryResponses(categories)
}
