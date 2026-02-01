package handler

import (
	"encoding/json"
	"go-pizza/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PizzaHandler struct {
	service  *service.PizzaService
	validate *validator.Validate
}

func NewPizzaHandler(s *service.PizzaService) *PizzaHandler {
	return &PizzaHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *PizzaHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		FlavorID string `json:"flavor_id" validate:"required"`
		Size     string `json:"size" validate:"required,oneof=P M G F"`
		ClientID string `json:"client_id" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(body); err != nil {

		validationErrors := err.(validator.ValidationErrors)

		var messages []string
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				messages = append(messages, e.Field()+" é obrigatório")
			case "oneof":
				messages = append(messages, e.Field()+" deve ser um dos seguintes valores: "+e.Param())
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"errors": messages,
		})
		return
	}

	order, err := h.service.CreateOrder(body.FlavorID, body.Size, body.ClientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go h.service.CookPizza(order.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *PizzaHandler) GetOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	order, err := h.service.GetOrderStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *PizzaHandler) GetAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orders, err := h.service.GetAllOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
