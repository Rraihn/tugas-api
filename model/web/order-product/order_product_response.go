package order_product

type OrderProductResponse struct {
	Id        int    `validate:"required"`
	OrderId   int    `json:"orderId"`
	ProductId int    `json:"productId"`
	Qty       int    `json:"qty"`
	Price     string `json:"price"`
}
