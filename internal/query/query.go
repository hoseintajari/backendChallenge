package query

import (
	"AecProject/internal/Utility"
	"AecProject/internal/database"
	"AecProject/internal/models"
	"time"
)

var db = database.Connection()

func FindUsername(name string) bool {
	result := db.Where("user_name =? ", name).First(&models.User{})
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func FindUserId(userId int) bool {
	result := db.Where("user_id=?", userId).Find(&models.Wallet{})
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func CheckPassword(username string, password string) bool {
	var user models.User
	result := db.Where("user_name=?", username).Find(&user)
	if result.Error != nil {
		return false
	}
	return user.Password == Utility.HashPassword(password)
}

func GetRoles(username string) bool {
	var user models.User
	db.Where("user_name=?", username).First(&user)
	response := user.Roles
	return response
}

func DailyTransactionTotal() {
	var t models.Transaction
	for {
		s := time.Now()
		t.Calculate(s)
		time.Sleep(24 * time.Hour)
	}
}
