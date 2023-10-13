package customers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Customer struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var customers []Customer

func Init() {
	customers = append(
		customers,
		Customer{
			Id:        "1",
			Name:      "John Doe",
			Role:      "Software Developer",
			Email:     "john.doe@gmail.com",
			Phone:     "1234567890",
			Contacted: true,
		},
	)
	customers = append(
		customers,
		Customer{
			Id:        "2",
			Name:      "Jane Doe",
			Role:      "Digital Marketer",
			Email:     "jane.doe@gmail.com",
			Phone:     "1234567890",
			Contacted: false,
		},
	)
	customers = append(
		customers,
		Customer{
			Id:        "3",
			Name:      "Terry Doe",
			Role:      "User Experience Designer",
			Email:     "terry.doe@gmail.com",
			Phone:     "1234567890",
			Contacted: true,
		},
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
