package product

type ProductResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
