package domain

import "github.com/google/uuid"

type OrderItem struct {
	id        int
	orderID   uuid.UUID
	productID uuid.UUID
	quantity  int
	price     float64
	subTotal  float64
}

func newOrderItem(ID int, orderID uuid.UUID, productID uuid.UUID, quantity int, price, subTotal float64) *OrderItem {
	return &OrderItem{
		id:        ID,
		orderID:   orderID,
		productID: productID,
		quantity:  quantity,
		price:     price,
		subTotal:  subTotal,
	}
}

func (o *OrderItem) SetProductID(productID uuid.UUID) {
	o.productID = productID
}

func (o *OrderItem) SetQuantity(quantity int) {
	o.quantity = quantity
}

func (o *OrderItem) SetPrice(price float64) {
	o.price = price
}

func (o *OrderItem) SetSubTotal(subTotal float64) {
	o.subTotal = subTotal
}
