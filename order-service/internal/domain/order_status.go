package domain

type OrderStatus string

const (
	Pending    OrderStatus = "PENDING"
	Paid       OrderStatus = "PAID"
	Approved   OrderStatus = "APPROVED"
	Cancelling OrderStatus = "CANCELLING"
	Cancelled  OrderStatus = "CANCELLED"
)
