package domain

type OrderStatus string

const (
	PENDING    OrderStatus = "PENDING"
	PAID       OrderStatus = "PAID"
	APPROVED   OrderStatus = "APPROVED"
	CANCELLING OrderStatus = "CANCELLING"
	CANCELLED  OrderStatus = "CANCELLED"
)
