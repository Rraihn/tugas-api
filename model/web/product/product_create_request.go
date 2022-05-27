package product

type ProductCreateRequest struct {
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Price      string `validate:"required,min=1,max=100" json:"price"`
	CategoryId int    `validate:"required,min=1,max=100" json:"categoryId"`
}
