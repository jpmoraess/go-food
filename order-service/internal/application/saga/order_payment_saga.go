package saga

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"log"
)

type OrderPaymentSaga struct {
}

func NewOrderPaymentSaga() *OrderPaymentSaga {
	return &OrderPaymentSaga{}
}

func (s *OrderPaymentSaga) Process(response *dto.PaymentResponse) error {
	log.Println("processing orderPaymentSaga")
	return nil
}

func (s *OrderPaymentSaga) Rollback(response *dto.PaymentResponse) error {
	log.Println("rollback orderPaymentSaga")
	return nil
}
