package customer

type CustomerResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
