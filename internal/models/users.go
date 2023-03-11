package models

import (
	"AecProject/internal/Utility"
	"AecProject/internal/controller/apimodels"
	"AecProject/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	UserID   int
	Password string
	Roles    bool `gorm:"default:false"`
	Email    string
}

var db = database.Connection()

func CreateUserTable() User {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Failed to Create Users table!")
	}
	if db.Where("user_name=?", "Admin").Find(&User{}).RowsAffected == 0 {
		db.Save(&User{UserName: "Admin", Password: Utility.HashPassword("admin"), Roles: true})
	}
	return User{}
}

func (u *User) Add(input apimodels.SignInRequest) {
	db.Save(&User{
		UserName: input.UserName,
		Password: Utility.HashPassword(input.Password),
		Email:    input.Email,
	})

}

func (u *User) Update() {

}

func (u *User) Remove(request apimodels.RemoveUserRequest) {
	db.Where("user_name=?", request.UserName).Find(&User{}).Delete(&User{})
}

func (u *User) RolesCheck(username string) bool {
	var s User
	db.Where("user_name =?", username).Find(&s)
	if s.Roles == true {
		return true
	}
	return false
}
