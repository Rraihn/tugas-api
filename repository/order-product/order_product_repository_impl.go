package order_product

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() OrderProductRepository {
	return &OrderProductRepositoryImpl{}
}

func (o OrderProductRepositoryImpl) SaveOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_product(order_id, product_id, qty, price) values (? ,? ,? ,?)"
	result, err := tx.ExecContext(ctx, SQL, orderproduct.OrderId, orderproduct.ProductId, orderproduct.Qty, orderproduct.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orderproduct.Id = int(id)
	return orderproduct
}

func (o OrderProductRepositoryImpl) UpdateOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) domain.OrderProduct {
	SQL := "update order_product set order_id = ?, product_id = ?, qty = ?, price = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderproduct.OrderId, orderproduct.ProductId, orderproduct.Qty, orderproduct.Price, orderproduct.Id)
	helper.PanicIfError(err)

	return orderproduct
}

func (o OrderProductRepositoryImpl) DeleteOrderProduct(ctx context.Context, tx *sql.Tx, orderproduct domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orderproduct.Id)
	helper.PanicIfError(err)
}

func (o OrderProductRepositoryImpl) FindOrderProductById(ctx context.Context, tx *sql.Tx, orderproductId int) (domain.OrderProduct, error) {
	SQL := "select id, order_id, product_id, qty, price from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, orderproductId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderProduct := domain.OrderProduct{}

	if rows.Next() {
		err := rows.Scan(&orderProduct.Id, &orderProduct.OrderId, &orderProduct.ProductId, &orderProduct.Qty, &orderProduct.Price)
		helper.PanicIfError(err)
		defer rows.Close()
		return orderProduct, nil
	} else {
		return orderProduct, errors.New("Order Product is not found")
	}
}

func (o OrderProductRepositoryImpl) FindAllOrderProduct(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	SQL := "select id, order_id, product_id, qty, price, created_at, updated_at from order_product "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderProducts []domain.OrderProduct
	for rows.Next() {
		orderProduct := domain.OrderProduct{}
		err := rows.Scan(&orderProduct.Id, &orderProduct.OrderId, &orderProduct.ProductId, &orderProduct.Qty, &orderProduct.Price, &orderProduct.CreatedAt, &orderProduct.UpdatedAt)
		helper.PanicIfError(err)
		orderProducts = append(orderProducts, orderProduct)
	}
	return orderProducts
}
