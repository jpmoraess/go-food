package domain

import (
	"fmt"
	"time"
)

type PaymentDomainService interface {
	ValidateAndInitiatePayment(payment *Payment, failureMessages *[]string) (PaymentEvent, error)
	ValidateAndCancelPayment(payment *Payment, failureMessages *[]string) (PaymentEvent, error)
}

type PaymentDomainServiceImpl struct{}

func NewPaymentDomainServiceImpl() *PaymentDomainServiceImpl {
	return &PaymentDomainServiceImpl{}
}

func (p *PaymentDomainServiceImpl) ValidateAndInitiatePayment(payment *Payment, failureMessages *[]string) (PaymentEvent, error) {
	payment.validatePayment(failureMessages)
	if len(*failureMessages) == 0 {
		fmt.Printf("Payment is initiated for order id: %s\n", payment.orderID.String())
		payment.updateStatus("COMPLETED")
		return newPaymentCompletedEvent(payment, time.Now()), nil
	} else {
		fmt.Printf("Payment initiation is failed for order id: %s\n", payment.orderID.String())
		payment.updateStatus("FAILED")
		return newPaymentFailedEvent(payment, time.Now(), *failureMessages), nil
	}
}

func (p *PaymentDomainServiceImpl) ValidateAndCancelPayment(payment *Payment, failureMessages *[]string) (PaymentEvent, error) {
	payment.validatePayment(failureMessages)
	if len(*failureMessages) == 0 {
		fmt.Printf("Payment is cancelled for order id: %s\n", payment.orderID.String())
		payment.updateStatus("CANCELLED")
		return newPaymentCancelledEvent(payment, time.Now()), nil
	} else {
		fmt.Printf("Payment cancellation is failed for order id: %s\n", payment.orderID.String())
		payment.updateStatus("FAILED")
		return newPaymentFailedEvent(payment, time.Now(), *failureMessages), nil
	}
}
