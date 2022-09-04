package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func CreateMember(w http.ResponseWriter, r *http.Request) {
	dsn := "root:root@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(w, "can't connect to db")
	}

	l := Login{
		Email:        "test@test.com",
		PasswordHash: "aaaaa",
		FirstName:    "Ftest",
		LastName:     "Ltest",
	}

	db.Create(&l)

	fmt.Fprintf(w, "created record: %v", l)

}
