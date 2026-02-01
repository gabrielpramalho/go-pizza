package repository

import (
	"errors"
	"go-pizza/internal/entity"
	"time"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	GetByID(id string) (entity.Order, error)
	UpdateStatus(id string, status entity.OrderStatus) (entity.Order, error)
	FindAll() ([]entity.Order, error)
}

type MemoryOrderRepository struct {
	orders map[string]entity.Order
}

func NewMemoryRepository() *MemoryOrderRepository {
	orders := make(map[string]entity.Order)
	return &MemoryOrderRepository{orders: orders}
}

func (r *MemoryOrderRepository) Create(order entity.Order) (entity.Order, error) {
	r.orders[order.ID] = entity.Order{
		ID:         order.ID,
		Size:       order.Size,
		FlavorID:   order.FlavorID,
		ClientID:   order.ClientID,
		TotalPrice: order.TotalPrice,
		Status:     "created",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return r.orders[order.ID], nil
}

func (r *MemoryOrderRepository) GetByID(id string) (entity.Order, error) {
	order, exists := r.orders[id]

	if !exists {
		return entity.Order{}, errors.New("Order not found")
	}
	return order, nil
}

func (r *MemoryOrderRepository) UpdateStatus(id string, status entity.OrderStatus) (entity.Order, error) {
	order, exists := r.orders[id]

	if !exists {
		return entity.Order{}, errors.New("Order not found")
	}

	order.Status = status
	order.UpdatedAt = time.Now()
	r.orders[id] = order

	return order, nil
}
