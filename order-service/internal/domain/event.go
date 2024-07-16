package domain

import "time"

type OrderEvent struct {
	Order     *Order
	CreatedAt time.Time
}

func (e *OrderEvent) GetOrder() *Order {
	return e.Order
}

func (e *OrderEvent) GetCreatedAt() time.Time {
	return e.CreatedAt
}

// OrderCreatedEvent -
type OrderCreatedEvent struct {
	OrderEvent
}

func NewOrderCreatedEvent(order *Order, createdAt time.Time) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		OrderEvent{
			Order:     order,
			CreatedAt: createdAt,
		},
	}
}

// OrderPaidEvent -
type OrderPaidEvent struct {
	OrderEvent
}

func NewOrderPaidEvent(order *Order, createdAt time.Time) *OrderPaidEvent {
	return &OrderPaidEvent{
		OrderEvent{
			Order:     order,
			CreatedAt: createdAt,
		},
	}
}

// OrderCancelledEvent -
type OrderCancelledEvent struct {
	OrderEvent
}

func NewOrderCancelledEvent(order *Order, createdAt time.Time) *OrderCancelledEvent {
	return &OrderCancelledEvent{
		OrderEvent{
			Order:     order,
			CreatedAt: createdAt,
		},
	}
}
