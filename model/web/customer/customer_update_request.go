package customer

import "time"

type CustomerUpdateRequest struct {
	Id          int       `validate:"required"`
	Name        string    `validate:"required,max=200,min=1" json:"name"`
	Email       string    `validate:"required,min=1,max=100" json:"email"`
	Address     string    `validate:"required,min=1,max=100" json:"address"`
	PhoneNumber string    `validate:"required,min=1,max=100" json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
