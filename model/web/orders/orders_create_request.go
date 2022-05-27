package orders

type OrdersCreateRequest struct {
	CustomerId int `validate:"required,min=1,max=100" json:"CustomerId"`
}
