package service

import (
	"fmt"
	"go-pizza/internal/entity"
	"go-pizza/internal/repository"
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

func (s *PizzaService) CreateOrder(flavorID, size, clientID string) (entity.Order, error) {
	price := float64(rand.Intn(50) + 20)

	order := entity.Order{
		ID:         uuid.New().String(),
		FlavorID:   flavorID,
		Size:       size,
		ClientID:   clientID,
		TotalPrice: price,
		Status:     "created",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return s.repo.Create(order)
}

func (s *PizzaService) CookPizza(id string) {
	fmt.Printf("🔥 [Cozinha] Pizza %s começou a assar...\n", id)
	time.Sleep(20 * time.Second)
	s.repo.UpdateStatus(id, "ready")
	fmt.Printf("🔔 [Cozinha] Pizza %s está PRONTA!\n", id)
}

func (s *PizzaService) GetOrderStatus(id string) (entity.Order, error) {
	return s.repo.GetByID(id)
}

func (s *PizzaService) GetAllOrders() ([]entity.Order, error) {
	return s.repo.FindAll()
}