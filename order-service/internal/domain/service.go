package domain

import (
	"log"
	"time"
)

type OrderDomainService interface {
	ValidateAndInitiateOrder(order *Order) (*OrderCreatedEvent, error)
	PayOrder(order *Order) (*OrderPaidEvent, error)
	CancelOrder(order *Order, failureMessages []string) (*OrderCancelledEvent, error)
}

type OrderDomainServiceImpl struct {
}

func (s *OrderDomainServiceImpl) ValidateAndInitiateOrder(order *Order) (*OrderCreatedEvent, error) {
	err := order.Validate()
	if err != nil {
		log.Printf("order validate failed: %v", err)
		return nil, err
	}
	return NewOrderCreatedEvent(order, time.Now()), nil
}

func (s *OrderDomainServiceImpl) PayOrder(order *Order) (*OrderPaidEvent, error) {
	err := order.Pay()
	if err != nil {
		log.Printf("order paid failed: %v", err)
		return nil, err
	}
	return NewOrderPaidEvent(order, time.Now()), nil
}

func (s *OrderDomainServiceImpl) CancelOrder(order *Order, failureMessages []string) (*OrderCancelledEvent, error) {
	err := order.Cancel(failureMessages)
	if err != nil {
		log.Printf("order cancel failed: %v", err)
		return nil, err
	}
	return NewOrderCancelledEvent(order, time.Now()), nil
}
