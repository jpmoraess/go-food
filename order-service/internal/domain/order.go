package domain

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrPriceGreaterThanZero                 = errors.New("price must greater than zero")
	ErrIncorrectStateForPayOperation        = errors.New("order is not in correct state for pay operation")
	ErrIncorrectStateForCancelOperation     = errors.New("order is not in correct state for cancel operation")
	ErrIncorrectStateForApproveOperation    = errors.New("order is not in correct state for approve operation")
	ErrIncorrectStateForInitCancelOperation = errors.New("order is not in correct state for init cancel operation")
)

type Order struct {
	id              uuid.UUID
	customerID      uuid.UUID
	restaurantID    uuid.UUID
	price           float64
	items           []*OrderItem
	orderStatus     OrderStatus
	trackingID      uuid.UUID
	failureMessages []string
}

func newOrder(customerID uuid.UUID, restaurantID uuid.UUID, price float64, items []*OrderItem) (*Order, error) {
	orderID, _ := uuid.NewV7()
	trackingID, _ := uuid.NewV7()

	if !(price > 0) {
		return nil, ErrPriceGreaterThanZero
	}

	return &Order{
		id:              orderID,
		customerID:      customerID,
		restaurantID:    restaurantID,
		price:           price,
		items:           items,
		orderStatus:     PENDING,
		trackingID:      trackingID,
		failureMessages: make([]string, 0),
	}, nil
}

func (o *Order) pay() error {
	if o.orderStatus != PENDING {
		return ErrIncorrectStateForPayOperation
	}
	o.orderStatus = PAID
	return nil
}

func (o *Order) approve() error {
	if o.orderStatus != PAID {
		return ErrIncorrectStateForApproveOperation
	}
	o.orderStatus = APPROVED
	return nil
}

func (o *Order) initCancel(failureMessages []string) error {
	if o.orderStatus != PAID {
		return ErrIncorrectStateForInitCancelOperation
	}
	o.orderStatus = CANCELLING
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) cancel(failureMessages []string) error {
	if !(o.orderStatus == CANCELLING || o.orderStatus == PENDING) {
		return ErrIncorrectStateForCancelOperation
	}
	o.orderStatus = CANCELLED
	o.updateFailureMessages(failureMessages)
	return nil
}

func (o *Order) updateFailureMessages(failureMessages []string) {
	if o.failureMessages != nil && failureMessages != nil {
		o.failureMessages = append(o.failureMessages, failureMessages...)
	}
	if o.failureMessages == nil {
		o.failureMessages = failureMessages
	}
}

func (o *Order) ID() uuid.UUID {
	return o.id
}

func (o *Order) CustomerID() uuid.UUID {
	return o.customerID
}

func (o *Order) SetCustomerID(customerID uuid.UUID) {
	o.customerID = customerID
}

func (o *Order) RestaurantID() uuid.UUID {
	return o.restaurantID
}

func (o *Order) SetRestaurantID(restaurantID uuid.UUID) {
	o.restaurantID = restaurantID
}

func (o *Order) Price() float64 {
	return o.price
}

func (o *Order) SetPrice(price float64) {
	o.price = price
}

func (o *Order) Items() []*OrderItem {
	return o.items
}

func (o *Order) SetItems(items []*OrderItem) {
	o.items = items
}

func (o *Order) Status() OrderStatus {
	return o.orderStatus
}

func (o *Order) TrackingID() uuid.UUID {
	return o.trackingID
}

func (o *Order) FailureMessages() []string {
	return o.failureMessages
}
