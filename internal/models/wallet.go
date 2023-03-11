package models

import (
	"AecProject/internal/controller/apimodels"
	"fmt"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserID  int
	Balance int
}

func CreateWalletTable() Wallet {
	err := db.AutoMigrate(&Wallet{})
	if err != nil {
		panic("Failed to create table")
	}
	return Wallet{}
}

func (w *Wallet) Add(request apimodels.NewWalletRequest) {
	tx := db.Save(&Wallet{UserID: request.UserID, Balance: request.Balance})
	fmt.Println(tx)
	db.Save(w.UserID)
}

func (w *Wallet) AddMoney(request apimodels.AddMoneyRequest) {
	var s Wallet
	db.Where("user_id =?", request.UserID).First(&s)
	s.Balance += request.Amount
	db.Save(&s)
}

func (w *Wallet) GetBalance(request apimodels.GetBalanceRequest) int {
	var s Wallet
	db.Where("user_id = ?", request.UserId).First(&s)
	response := s.Balance
	return response
}
