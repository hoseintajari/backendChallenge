package main

import (
	config2 "AecProject/internal/config"
	"AecProject/internal/models"
	"AecProject/internal/query"
	"AecProject/internal/routing"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	models.CreateUserTable()
	models.CreateWalletTable()
	models.CreateTransactionTable()

	config, err := config2.LoadConfig("./cmd")
	if err != nil {
		panic(err)
	}
	serverAddress := fmt.Sprintf("%s:%d", config.ServerIp, config.ServerPort)

	server := echo.New()
	routing.MainRout(server)
	go func() { query.DailyTransactionTotal() }()
	if err := server.Start(serverAddress); err != nil {
		server.Logger.Info("shutting down the server")
	}
}
