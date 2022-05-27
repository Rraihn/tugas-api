package orders

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (o OrdersRepositoryImpl) SaveOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "insert into orders(customer_id) values (?)"
	result, err := tx.ExecContext(ctx, SQL, orders.CustomerID)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orders.Id = int(id)
	return orders
}

func (o OrdersRepositoryImpl) UpdateOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "update orders set customer_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.CustomerID, orders.Id)
	helper.PanicIfError(err)

	return orders
}

func (o OrdersRepositoryImpl) DeleteOrders(ctx context.Context, tx *sql.Tx, orders domain.Orders) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.Id)
	helper.PanicIfError(err)
}

func (o OrdersRepositoryImpl) FindOrdersById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error) {
	SQL := "select id, customer_id from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ordersId)
	helper.PanicIfError(err)
	defer rows.Close()

	orders := domain.Orders{}

	if rows.Next() {
		err := rows.Scan(&orders.Id, &orders.CustomerID)
		helper.PanicIfError(err)
		defer rows.Close()
		return orders, nil
	} else {
		return orders, errors.New("Category is not found")
	}
}

func (o OrdersRepositoryImpl) FindAllOrders(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "select id, customer_id from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var order []domain.Orders
	for rows.Next() {
		orders := domain.Orders{}
		err := rows.Scan(&orders.Id, &orders.CustomerID)
		helper.PanicIfError(err)
		order = append(order, orders)
	}
	return order
}
