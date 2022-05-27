package product

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required,max=200,min=1" json:"name"`
	Price      string `validate:"required,max=200,min=1" json:"price"`
	CategoryId int    `validate:"required,max=200,min=1" json:"categoryId"`
}
