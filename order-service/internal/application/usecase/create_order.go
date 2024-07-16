package usecase

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/helper"
)

type CreateOrderUseCase struct {
	helper helper.CreateOrderHelper
}

func NewCreateOrderUseCase(helper helper.CreateOrderHelper) *CreateOrderUseCase {
	return &CreateOrderUseCase{helper: helper}
}

func (uc *CreateOrderUseCase) Execute(input *dto.CreateOrderInputDTO) (*dto.CreateOrderOutputDTO, error) {
	// TODO: PersistOrder and OutboxMessage should be a single transaction scope UnitOfWork
	event, err := uc.helper.PersistOrder(input)
	if err != nil {
		return nil, err
	}

	// TODO: save outbox message

	return &dto.CreateOrderOutputDTO{
		TrackingID: event.Order.TrackingID,
		Message:    "order created successfully",
	}, nil
}
