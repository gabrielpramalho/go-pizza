package service

import (
	"go-pizza/internal/entity"
	"go-pizza/internal/repository"
	"log/slog"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type PizzaService struct {
	repo repository.OrderRepository
}

func NewPizzaService(repo repository.OrderRepository) *PizzaService {
	return &PizzaService{repo: repo}
}

func (s *PizzaService) CreateOrder(flavorID, clientID string, size entity.PizzaSize) (entity.Order, error) {
	price := float64(rand.Intn(50) + 20)

	order := entity.Order{
		ID:         uuid.New().String(),
		FlavorID:   flavorID,
		Size:       size,
		ClientID:   clientID,
		TotalPrice: price,
		Status:     entity.OrderStatusPending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.repo.Create(order)
}

func (s *PizzaService) CookPizza(id string) {
	slog.Info("Pizza começou a assar", "order_id", id)
	time.Sleep(20 * time.Second)
	order, err := s.repo.UpdateStatus(id, entity.OrderStatusReady)
	if err != nil {
		slog.Error("Erro ao atualizar status da pizza", "order_id", id, "error", err)
		return
	}
	slog.Info("Pizza pronta", "order_id", id, "client_id", order.ClientID)
}

func (s *PizzaService) GetOrderStatus(id string) (entity.Order, error) {
	return s.repo.GetByID(id)
}

func (s *PizzaService) GetAllOrders() ([]entity.Order, error) {
	return s.repo.FindAll()
}
