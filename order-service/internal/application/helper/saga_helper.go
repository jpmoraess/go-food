package helper

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/enum"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type SagaHelper struct{}

func (s *SagaHelper) OrderStatusToSagaStatus(orderStatus domain.OrderStatus) enum.SagaStatus {
	switch orderStatus {
	case domain.PAID:
		return enum.SAGA_PROCESSING
	case domain.APPROVED:
		return enum.SAGA_SUCCEEDED
	case domain.CANCELLING:
		return enum.SAGA_COMPENSATING
	case domain.CANCELLED:
		return enum.SAGA_COMPENSATED
	default:
		return enum.SAGA_STARTED
	}
}
