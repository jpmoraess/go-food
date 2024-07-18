package usecase

import (
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
)

type CreateOrderUseCase struct {
	orderSagaHelper     helper.SagaHelper
	createOrderHelper   helper.CreateOrderHelper
	paymentOutboxHelper outbox.PaymentOutboxHelper
}

func NewCreateOrderUseCase(orderSagaHelper helper.SagaHelper, createOrderHelper helper.CreateOrderHelper, paymentOutboxHelper outbox.PaymentOutboxHelper) *CreateOrderUseCase {
	return &CreateOrderUseCase{orderSagaHelper: orderSagaHelper, createOrderHelper: createOrderHelper, paymentOutboxHelper: paymentOutboxHelper}
}

func (uc *CreateOrderUseCase) Execute(input *dto.CreateOrderInputDTO) (*dto.CreateOrderOutputDTO, error) {
	event, err := uc.createOrderHelper.PersistOrder(input)
	if err != nil {
		return nil, err
	}

	sagaId, _ := uuid.NewRandom()
	err = uc.paymentOutboxHelper.SavePaymentOutbox(
		nil,
		event.Order.OrderStatus,
		uc.orderSagaHelper.OrderStatusToSagaStatus(event.Order.OrderStatus),
		outbox.STARTED,
		sagaId,
	)
	if err != nil {
		return nil, err
	}

	return &dto.CreateOrderOutputDTO{
		TrackingID: event.Order.TrackingID,
		Message:    "order created successfully",
	}, nil
}
