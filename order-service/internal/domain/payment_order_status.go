package domain

type PaymentOrderStatus string

const (
	PAYMENT_PENDING   PaymentOrderStatus = "PENDING"
	PAYMENT_CANCELLED PaymentOrderStatus = "CANCELLED"
)
