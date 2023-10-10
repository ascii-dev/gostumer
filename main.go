package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var customers map[string]Customer

func main() {
	router := mux.NewRouter()

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}
