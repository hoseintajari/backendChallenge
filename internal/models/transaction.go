package models

import (
	"AecProject/internal/controller/apimodels"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type Transaction struct {
	gorm.Model
	UserID      int
	Amount      int
	ReferenceID int
}

func CreateTransactionTable() Transaction {

	err := db.AutoMigrate(&Transaction{})
	if err != nil {
		panic("Failed to create table")
	}
	return Transaction{}
}

func (t *Transaction) Add(request apimodels.AddMoneyRequest) int {
	response := rand.Int()

	db.Save(&Transaction{UserID: request.UserID, Amount: request.Amount, ReferenceID: response})
	return response
}

func (t *Transaction) Calculate(duration time.Time) {
	var total float64
	if err := db.Model(&Transaction{}).Select("COALESCE(SUM(amount), 0) as total").Where("created_at >= ?", duration).Scan(&total).Error; err != nil {
		log.Fatal(err)
	}

	log.Println("Total transactions in last 24 hours:", total)
}
