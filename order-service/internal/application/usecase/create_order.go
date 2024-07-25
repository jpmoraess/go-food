package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
	"github.com/jpmoraess/go-food/order-service/internal/application/mapper"
)

type CreateOrderUseCase struct {
	orderMapper         *mapper.OrderMapper
	orderSagaHelper     *helper.SagaHelper
	createOrderHelper   *helper.CreateOrderHelper
	paymentOutboxHelper *helper.PaymentOutboxHelper
}

func NewCreateOrderUseCase(orderMapper *mapper.OrderMapper, orderSagaHelper *helper.SagaHelper, createOrderHelper *helper.CreateOrderHelper, paymentOutboxHelper *helper.PaymentOutboxHelper) *CreateOrderUseCase {
	return &CreateOrderUseCase{orderMapper: orderMapper, orderSagaHelper: orderSagaHelper, createOrderHelper: createOrderHelper, paymentOutboxHelper: paymentOutboxHelper}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input *dto.CreateOrderInputDTO) (*dto.CreateOrderOutputDTO, error) {
	orderCreatedEvent, err := uc.createOrderHelper.PersistOrder(ctx, input)
	if err != nil {
		return nil, err
	}

	err = uc.paymentOutboxHelper.SavePaymentOutbox(
		ctx,
		uc.orderMapper.OrderCreatedEventToPaymentEventPayload(orderCreatedEvent),
		orderCreatedEvent.Order().Status(),
		uc.orderSagaHelper.OrderStatusToSagaStatus(orderCreatedEvent.Order().Status()),
		enum.OUTBOX_STARTED,
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
