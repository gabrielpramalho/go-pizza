package main

import (
	"go-pizza/internal/handler"
	"go-pizza/internal/repository"
	"go-pizza/internal/service"
	"log/slog"
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

	slog.Info("Conectando ao Postgres...")
	repo, err := repository.NewPostgresRepository(dbURL)
	if err != nil {
		slog.Error("Erro ao conectar ao Postgres", "error", err)
		panic(err)
	}
	svc := service.NewPizzaService(repo)
	hand := handler.NewPizzaHandler(svc)

	http.HandleFunc("/orders", hand.CreateOrderHandler)
	http.HandleFunc("/orders/status", hand.GetOrderStatusHandler)
	http.HandleFunc("/orders/all", hand.GetAllOrdersHandler)

	slog.Info("Servidor iniciado", "port", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Error("Erro ao iniciar o servidor", "error", err)
		panic(err)
	}
}
