package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web/product"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) SaveProduct(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("Product Repository Save Product Start")
	SQL := "insert into product(name, price, category_id) values (?, ? ,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	logrus.Info("Product Repository SaveProduct End")
	return product
}

func (p ProductRepositoryImpl) UpdateProduct(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("Product Repository Update Product Start")
	SQL := "update product set name = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	logrus.Info("Product Repository Update Product End")
	return product
}

func (p ProductRepositoryImpl) DeleteProduct(ctx context.Context, tx *sql.Tx, product domain.Product) {
	logrus.Info("Product Repository Delete Product Start")
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
	logrus.Info("Product Repository Delete Product End")
}

func (p ProductRepositoryImpl) FindProductById(ctx context.Context, tx *sql.Tx, productId int) (product.ProductResponse, error) {
	logrus.Info("Product Repository Find Product By Id Start")
	SQL := "select p.id, p.name, p.price, p.category_id, c.name as 'category_name' from product p inner join category c on p.category_id=c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := product.ProductResponse{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		defer rows.Close()
		logrus.Info("Product Repository Find Product By Id End")
		return product, nil
	} else {
		logrus.Info("Product Repository Find Product By Id End")
		return product, errors.New("Product is not found")
	}
}

func (p ProductRepositoryImpl) FindAllProduct(ctx context.Context, tx *sql.Tx) []product.ProductResponse {
	logrus.Info("Product Repository Find All Product Start")
	SQL := "select p.id, p.name, p.price, p.category_id, c.name as 'category_name' from product p inner join category c on p.category_id=c.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []product.ProductResponse
	for rows.Next() {
		product := product.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	logrus.Info("Product Repository Find All Product End")
	return products
}
