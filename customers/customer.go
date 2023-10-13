package customers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Customer struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var customers []Customer

func Init() {
	customers = append(
		customers,
		Customer{Id: "1", FirstName: "John", LastName: "Doe", Email: "john.doe@gmail.com", Phone: "1234567890"},
	)
	customers = append(
		customers,
		Customer{Id: "2", FirstName: "Jane", LastName: "Doe", Email: "jane.doe@gmail.com", Phone: "1234567890"},
	)
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}

func GetSingleCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid customer ID"})
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(customers[id-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid customer ID"})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		return
	}

	customer.Id = strconv.Itoa(len(customers) + 1)
	customers = append(customers, customer)

	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid customer ID"})
		if err != nil {
			return
		}
		return
	}

	customer.Id = strconv.Itoa(id)
	customers[id-1] = customer

	err = json.NewEncoder(w).Encode(customers)
	w.WriteHeader(http.StatusOK)
	if err != nil {
		return
	}
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid customer ID"})
		if err != nil {
			return
		}
		return
	}

	customers = append(customers[:id-1], customers[id:]...)

	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}
