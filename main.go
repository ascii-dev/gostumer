package main

import (
	"fmt"
	"github.com/ascii-dev/gostumer/customers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	customers.Init()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/customers", customers.GetCustomers).Methods("GET")
	router.HandleFunc("/api/v1/customers", customers.CreateCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customers/{id}", customers.GetSingleCustomer).Methods("GET")
	router.HandleFunc("/api/v1/customers/{id}", customers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/api/v1/customers/{id}", customers.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		return
	}
}
