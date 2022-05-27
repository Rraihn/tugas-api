package orders

type OrdersUpdateRequest struct {
	Id         int `validate:"required"`
	CustomerId int `validate:"required,max=200,min=1" json:"CustomerId"`
}
