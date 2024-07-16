package domain

import "github.com/google/uuid"

type OrderItem struct {
	ID       int32
	OrderID  uuid.UUID
	Product  *Product
	Quantity int32
	Price    float64
	SubTotal float64
}

func NewOrderItem(orderID uuid.UUID, ID int32) *OrderItem {
	return &OrderItem{
		ID:      ID,
		OrderID: orderID,
	}
}

func (o *OrderItem) IsPriceValid() bool {
	return o.Price > 0 && o.Price == o.Product.Price && (o.Price*float64(o.Quantity)) == o.SubTotal
}
