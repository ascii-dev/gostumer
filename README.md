# gostumer

Gostumer is a Customer Relationship Management API that handles the information
of onboarded customers.

## Run The Application
To run the application, you need to have at least go.19 installed.

- Open the project in your favorite editor
- Run `go mod download` to install dependencies
- Start the application using `go run main.go`
- Make requests to the endpoints through Postman or cURL

## Endpoints
- Get all customers: `GET /api/v1/customers`
- Get a single customer: `GET /api/v1/customers/{id}`
- Update a single customer: `PUT /api/v1/customers/{id}`
- Add a new customer: `POST /api/v1/customers`
- Delete a customer: `DELETE /api/v1/customers/{id}`
