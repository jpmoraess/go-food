package dto

import "github.com/google/uuid"

type CreateOrderInputDTO struct {
	CustomerID   uuid.UUID            `json:"customerId"`
	RestaurantID uuid.UUID            `json:"restaurantId"`
	Price        float64              `json:"price"`
	Items        []OrderItemInputDTO  `json:"items"`
	Address      OrderAddressInputDTO `json:"address"`
}

type CreateOrderOutputDTO struct {
	TrackingID uuid.UUID `json:"trackingId"`
	Message    string    `json:"message"`
}

type OrderItemInputDTO struct {
	ProductID uuid.UUID `json:"productId"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	SubTotal  float64   `json:"subTotal"`
}

type OrderAddressInputDTO struct {
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
}
