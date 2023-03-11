package main

import (
	"AecProject/internal/models"
	"AecProject/internal/query"
	"AecProject/internal/routing"
	"github.com/labstack/echo/v4"
)

func main() {
	models.CreateUserTable()
	models.CreateWalletTable()
	models.CreateTransactionTable()
	server := echo.New()
	routing.MainRout(server)
	go func() { query.DailyTransactionTotal() }()
	if err := server.Start(":8081"); err != nil {
		server.Logger.Info("shutting down the server")
	}
}
