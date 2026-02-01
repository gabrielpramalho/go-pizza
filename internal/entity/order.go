package entity

import "time"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusCooking   OrderStatus = "COOKING"
	OrderStatusReady     OrderStatus = "READY"
	OrderStatusDelivered OrderStatus = "DELIVERED"
)

type PizzaSize string

const (
	PizzaSizeP PizzaSize = "P"
	PizzaSizeM PizzaSize = "M"
	PizzaSizeG PizzaSize = "G"
	PizzaSizeF PizzaSize = "F"
)

type Order struct {
	ID         string      `json:"id"`
	Size       PizzaSize   `json:"size"`
	FlavorID   string      `json:"flavor"`
	ClientID   string      `json:"client_id"`
	Status     OrderStatus `json:"status"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
