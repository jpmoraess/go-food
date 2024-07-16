package dto

import "github.com/google/uuid"

type CreateOrderInputDTO struct {
}

type CreateOrderOutputDTO struct {
	TrackingID uuid.UUID `json:"trackingId"`
	Message    string    `json:"message"`
}
