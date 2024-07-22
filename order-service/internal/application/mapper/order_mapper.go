package mapper

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type OrderMapper struct{}

func (o *OrderMapper) CreateOrderInputToOrder(input *dto.CreateOrderInputDTO) *domain.Order {
	order := new(domain.Order)
	order.SetCustomerID(input.CustomerID)
	order.SetRestaurantID(input.RestaurantID)
	order.SetPrice(input.Price)
	order.SetItems(o.orderItemsInputToOrderItems(input.Items))
	return order
}

func (o *OrderMapper) orderItemsInputToOrderItems(items []dto.OrderItemInputDTO) []*domain.OrderItem {
	orderItems := make([]*domain.OrderItem, len(items))
	for i, orderItemDTO := range items {
		orderItem := new(domain.OrderItem)
		orderItem.SetProductID(orderItemDTO.ProductID)
		orderItem.SetQuantity(orderItemDTO.Quantity)
		orderItem.SetPrice(orderItemDTO.Price)
		orderItem.SetSubTotal(orderItemDTO.SubTotal)
		orderItems[i] = orderItem
	}
	return orderItems
}

func (o *OrderMapper) OrderCreatedEventToPaymentEventPayload(event *domain.OrderCreatedEvent) *outbox.PaymentEventPayload {
	return &outbox.PaymentEventPayload{
		OrderID:            event.Order().ID().String(),
		CustomerID:         event.Order().CustomerID().String(),
		Price:              event.Order().Price(),
		CreatedAt:          event.CreatedAt(),
		PaymentOrderStatus: domain.PAYMENT_PENDING,
	}
}
