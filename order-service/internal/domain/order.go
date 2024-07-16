package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrPriceGreaterThanZero = errors.New("price must greater than zero")
)

type Order struct {
	ID              uuid.UUID
	CustomerID      uuid.UUID
	RestaurantID    uuid.UUID
	Price           float64
	Items           []*OrderItem
	OrderStatus     OrderStatus
	TrackingID      uuid.UUID
	FailureMessages []string
}

func NewOrder(customerID uuid.UUID, restaurantID uuid.UUID) *Order {
	orderID, err := uuid.NewV7()
	if err != nil {

	}
	trackingID, err := uuid.NewV7()
	if err != nil {

	}
	return &Order{
		ID:           orderID,
		CustomerID:   customerID,
		RestaurantID: restaurantID,
		TrackingID:   trackingID,
		OrderStatus:  Pending,
	}
}

func (o *Order) Validate() error {
	err := o.validateTotalPrice()
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) Pay() error {
	if o.OrderStatus != Pending {
		return errors.New("order is not in correct state for pay operation")
	}
	o.OrderStatus = Paid
	return nil
}

func (o *Order) Approve() error {
	if o.OrderStatus != Paid {
		return errors.New("order is not in correct state for approve operation")
	}
	o.OrderStatus = Approved
	return nil
}

func (o *Order) InitCancel(failureMessages []string) error {
	if o.OrderStatus != Paid {
		return errors.New("order is not in correct state for init cancel operation")
	}
	o.OrderStatus = Cancelling
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) Cancel(failureMessages []string) error {
	if !(o.OrderStatus == Cancelling || o.OrderStatus == Pending) {
		return errors.New("order is not in correct state for cancel operation")
	}
	o.OrderStatus = Cancelled
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) validateTotalPrice() error {
	if !(o.Price > 0) {
		return errors.New("total price must be greater than zero")
	}
	return nil
}

func (o *Order) initializeOrderItems() {
	for i := 0; i < len(o.Items); i++ {

	}
}

func (o *Order) updateFailureMessages(failureMessages []string) {
	if o.FailureMessages != nil && failureMessages != nil {
		o.FailureMessages = append(o.FailureMessages, failureMessages...)
	}
	if o.FailureMessages == nil {
		o.FailureMessages = failureMessages
	}
}
