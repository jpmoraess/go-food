package domain

import "time"

type orderEvent struct {
	order     *Order
	createdAt time.Time
}

func (e *orderEvent) Order() *Order {
	return e.order
}

func (e *orderEvent) CreatedAt() time.Time {
	return e.createdAt
}

// OrderCreatedEvent -
type OrderCreatedEvent struct {
	orderEvent
}

func newOrderCreatedEvent(order *Order, createdAt time.Time) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		orderEvent{
			order:     order,
			createdAt: createdAt,
		},
	}
}

// OrderPaidEvent -
type OrderPaidEvent struct {
	orderEvent
}

func newOrderPaidEvent(order *Order, createdAt time.Time) *OrderPaidEvent {
	return &OrderPaidEvent{
		orderEvent{
			order:     order,
			createdAt: createdAt,
		},
	}
}

// OrderCancelledEvent -
type OrderCancelledEvent struct {
	orderEvent
}

func newOrderCancelledEvent(order *Order, createdAt time.Time) *OrderCancelledEvent {
	return &OrderCancelledEvent{
		orderEvent{
			order:     order,
			createdAt: createdAt,
		},
	}
}
