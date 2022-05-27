package domain

import "time"

type OrderProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Qty       int
	Price     string
	Amount    string
	CreatedAt time.Time
	updatedAt time.Time
}
