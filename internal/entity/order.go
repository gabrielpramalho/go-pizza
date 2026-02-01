package entity

import "time"

type Order struct {
	ID         string    `json:"id"`
	Size       string    `json:"size"`
	FlavorID   string    `json:"flavor"`
	ClientID   string    `json:"client_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
