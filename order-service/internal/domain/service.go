package domain

import (
	"log"
	"time"
)

type OrderDomainService interface {
	InitiateOrder(order *Order) (*OrderCreatedEvent, error)
	PayOrder(order *Order) (*OrderPaidEvent, error)
	CancelOrder(order *Order, failureMessages []string) (*OrderCancelledEvent, error)
}

type OrderDomainServiceImpl struct {
}

func (s *OrderDomainServiceImpl) InitiateOrder(order *Order) (*OrderCreatedEvent, error) {
	entity, err := newOrder(order.CustomerID(), order.RestaurantID(), order.Price(), order.Items())
	if err != nil {
		log.Printf("create order failed: %v", err)
		return nil, err
	}
	return newOrderCreatedEvent(entity, time.Now()), nil
}

func (s *OrderDomainServiceImpl) PayOrder(order *Order) (*OrderPaidEvent, error) {
	err := order.pay()
	if err != nil {
		log.Printf("order paid failed: %v", err)
		return nil, err
	}
	return newOrderPaidEvent(order, time.Now()), nil
}

func (s *OrderDomainServiceImpl) CancelOrder(order *Order, failureMessages []string) (*OrderCancelledEvent, error) {
	err := order.cancel(failureMessages)
	if err != nil {
		log.Printf("order cancel failed: %v", err)
		return nil, err
	}
	return newOrderCancelledEvent(order, time.Now()), nil
}
