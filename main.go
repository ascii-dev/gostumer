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

	router.HandleFunc("/customers", customers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", customers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", customers.GetSingleCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", customers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", customers.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		return
	}
}
