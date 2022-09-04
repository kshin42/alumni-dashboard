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
	router.HandleFunc("/dbsetup", SetUpDB).Methods("GET")

	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/member", CreateMember).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequests()
}
