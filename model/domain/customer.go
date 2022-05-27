package domain

import "time"

type Customer struct {
	Id          int
	Name        string
	Email       string
	Address     string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
