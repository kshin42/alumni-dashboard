package main

import (
	"fmt"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Login struct {
	ID           uint
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func SetUpDB(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(w, "can't connect to db")
	}

	db.AutoMigrate(&Login{})
}
