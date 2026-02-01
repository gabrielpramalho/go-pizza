package handler

import (
	"encoding/json"
	"go-pizza/internal/entity"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PizzaServiceInterface interface {
	CreateOrder(flavorID, clientID string, size entity.PizzaSize) (entity.Order, error)
	CookPizza(id string)
	GetOrderStatus(id string) (entity.Order, error)
	GetAllOrders() ([]entity.Order, error)
}

type PizzaHandler struct {
	service  PizzaServiceInterface
	validate *validator.Validate
}

func NewPizzaHandler(s PizzaServiceInterface) *PizzaHandler {
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

		writeJSONResponse(w, http.StatusBadRequest, map[string]any{"errors": messages})

		if err != nil {
			slog.Error("failed to write response", "error", err)
		}
		return
	}

	order, err := h.service.CreateOrder(body.FlavorID, body.ClientID, entity.PizzaSize(body.Size))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go h.service.CookPizza(order.ID)

	writeJSONResponse(w, http.StatusCreated, order)
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

	writeJSONResponse(w, http.StatusOK, order)
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

	writeJSONResponse(w, http.StatusOK, orders)
}
