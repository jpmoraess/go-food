package saga

import (
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/domain"
	"log"
)

type OrderPaymentSaga struct {
}

func NewOrderPaymentSaga() *OrderPaymentSaga {
	return &OrderPaymentSaga{}
}

func (s *OrderPaymentSaga) Process(response *dto.PaymentResponse) error {
	//log.Println("processing orderPaymentSaga")
	//sagaID, _ := uuid.Parse(response.SagaID)
	//paymentOutbox := s.paymentOutboxHelper.GetPaymentOutboxMessageBySagaIdAndSagaStatus(context.Background(), sagaID, STARTED)
	//if paymentOutbox == nil {
	//	return nil
	//}
	return nil
}

func (s *OrderPaymentSaga) Rollback(response *dto.PaymentResponse) error {
	log.Println("rollback orderPaymentSaga")
	return nil
}

func (s *OrderPaymentSaga) completePayment(response *dto.PaymentResponse) (*domain.OrderPaidEvent, error) {
	//orderID, _ := uuid.Parse(response.OrderID)
	//order, err := s.sagaHelper.FindOrder(context.Background(), orderID)
	//if err != nil {
	//	return nil, err
	//}
	//orderPaidEvent, err := s.domainService.PayOrder(order)
	//if err != nil {
	//	return nil, err
	//}
	//_, err = s.orderRepo.Save(context.Background(), order)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}
