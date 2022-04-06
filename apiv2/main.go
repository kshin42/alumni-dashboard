package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/createMember", CreateMember)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
