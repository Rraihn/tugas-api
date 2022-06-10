package orders

import "time"

type OrdersResponse struct {
	Id         int       `json:"id"`
	CustomerId int       `json:"CustomerId"`
	DateTime   time.Time `json:"dateTime"`
}
