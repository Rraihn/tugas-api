package category

import "time"

type CategoryResponse struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	DateTime time.Time `json:"dateTime"`
}
