package domain

import "time"

type Orders struct {
	Id          int
	OrderDate   time.Time
	CustomerID  int
	TotalAmount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
