package main

type Customer struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type OrderRequest struct {
	CustomerID  int64  `json:"customerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type OrderResponse struct {
	ID          int64  `json:"id"`
	CustomerID  int64  `json:"customerId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
