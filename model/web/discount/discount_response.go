package discount

import "time"

type Discount struct {
	discount_id     int       `json:"discount_id"`
	discount_Status string    `json:"discount_Status"`
	discount_status string    `json:"discount_Status"`
	discount_amount float64   `json:"discount_Amount"`
	invoice_id      int       `json:"invoice_Id"`
	created_at      time.Time `json:"created_At"`
	updated_at      time.Time `json:"updated_At"`
}
