package main

import (
	"fmt"
	"go-pizza/internal/handler"
	"go-pizza/internal/repository"
	"go-pizza/internal/service"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	if dbURL == "" {
		panic("A variável DATABASE_URL é obrigatória!")
	}

	fmt.Println("🐘 Conectando ao Postgres...")
	repo, err := repository.NewPostgresRepository(dbURL)
	if err != nil {
		panic(err)
	}
	svc := service.NewPizzaService(repo)
	hand := handler.NewPizzaHandler(svc)

	http.HandleFunc("/orders", hand.CreateOrderHandler)
	http.HandleFunc("/orders/status", hand.GetOrderStatusHandler)
	http.HandleFunc("/orders/all", hand.GetAllOrdersHandler)

	fmt.Printf("🚀 Servidor iniciado na porta %s \n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
