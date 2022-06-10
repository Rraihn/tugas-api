package customer

import "time"

type CustomerResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	DateTime    time.Time `json:"dateTime"`
}
