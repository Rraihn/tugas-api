package domain

import "time"

type Discount struct {
	discount_id     int
	discount_Status string
	discount_status string
	discount_amount float64
	invoice_id      int
	created_at      time.Time
	updated_at      time.Time
}
