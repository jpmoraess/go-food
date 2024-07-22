package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
	"github.com/jpmoraess/go-food/order-service/internal/application/outbox"
)

type CreateOrderUseCase struct {
	orderSagaHelper     *helper.SagaHelper
	createOrderHelper   *helper.CreateOrderHelper
	paymentOutboxHelper *outbox.PaymentOutboxHelper
}

func NewCreateOrderUseCase(orderSagaHelper *helper.SagaHelper, createOrderHelper *helper.CreateOrderHelper, paymentOutboxHelper *outbox.PaymentOutboxHelper) *CreateOrderUseCase {
	return &CreateOrderUseCase{orderSagaHelper: orderSagaHelper, createOrderHelper: createOrderHelper, paymentOutboxHelper: paymentOutboxHelper}
}

// TODO: persist and payment outbox (db transaction)
func (uc *CreateOrderUseCase) Execute(ctx context.Context, input *dto.CreateOrderInputDTO) (*dto.CreateOrderOutputDTO, error) {
	orderCreatedEvent, err := uc.createOrderHelper.PersistOrder(ctx, input)
	if err != nil {
		return nil, err
	}

	orderMapper := mapper.OrderMapper{}
	err = uc.paymentOutboxHelper.SavePaymentOutbox(
		ctx,
		orderMapper.OrderCreatedEventToPaymentEventPayload(orderCreatedEvent),
		orderCreatedEvent.Order().Status(),
		uc.orderSagaHelper.OrderStatusToSagaStatus(orderCreatedEvent.Order().Status()),
		outbox.STARTED,
		uuid.New(),
	)
	if err != nil {
		return nil, err
	}

	return &dto.CreateOrderOutputDTO{
		TrackingID: orderCreatedEvent.Order().TrackingID(),
		Message:    "order created successfully",
	}, nil
}
