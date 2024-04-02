package tools

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type RunescapeAccount struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Members  bool   `json:"members"`
	Wealth   int    `json:"wealth"`
}

type Item struct {
	gorm.Model
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
	Account string `json:"account"`
}


func NewDatabase() {
	DB, err = gorm.Open(sqlite.Open("lennissa.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open db.")
	}

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&RunescapeAccount{})
	DB.AutoMigrate(&Item{})

	DB.Create(&Account{Username: "admin", Password: "admin"})
}
