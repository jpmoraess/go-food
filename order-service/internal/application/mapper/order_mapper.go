package mapper

import (
	"github.com/google/uuid"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
)

type OrderMapper struct{}

func (o *OrderMapper) MapOrderInputToDomain(input *dto.CreateOrderInputDTO) *domain.Order {
	random, _ := uuid.NewRandom()
	return domain.NewOrder(random, random)
}
