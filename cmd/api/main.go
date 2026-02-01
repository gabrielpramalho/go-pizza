package main

import (
	"fmt"
	"go-pizza/internal/handler"
	"go-pizza/internal/repository"
	"go-pizza/internal/service"
	"net/http"
)

func main() {
	repo, err := repository.NewSQLiteRepository()
    if err != nil {
        panic(err)
    }
	svc := service.NewPizzaService(repo)
	hand := handler.NewPizzaHandler(svc)

	http.HandleFunc("/orders", hand.CreateOrderHandler)
	http.HandleFunc("/orders/status", hand.GetOrderStatusHandler)
	http.HandleFunc("/orders/all", hand.GetAllOrdersHandler)

	fmt.Println("🚀 Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
