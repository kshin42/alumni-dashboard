package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Login struct {
	Id           int
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
}

func CreateMember(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "telam.db")
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	defer db.Close()

	var login Login
	db.Find(&login)
	json.NewEncoder(w).Encode(login)

}
