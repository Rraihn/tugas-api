package order_product

import "time"

type OrderProductCreateRequest struct {
	OrderId   int       `validate:"required,min=1,max=100" json:"orderId"`
	ProductId int       `validate:"required,min=1,max=100" json:"productId"`
	Qty       int       `validate:"required,min=1,max=100" json:"qty"`
	Price     string    `validate:"required,min=1,max=100" json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
