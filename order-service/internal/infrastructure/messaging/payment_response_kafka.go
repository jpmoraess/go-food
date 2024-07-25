package messaging

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/jpmoraess/go-food/order-service/internal/application/gateway"
)

type PaymentResponseKafka struct {
	listener      gateway.PaymentMessageListener
	consumerGroup sarama.ConsumerGroup
}

func NewPaymentResponseKafka(listener gateway.PaymentMessageListener) *PaymentResponseKafka {
	brokers := []string{"localhost:19092", "localhost:29092", "localhost:39092"}
	groupID := "order-service"

	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Panicf("error creating consumer group client: %v", err)
	}

	return &PaymentResponseKafka{listener: listener, consumerGroup: consumerGroup}
}

type paymentConsumerGroupHandler struct {
	listener gateway.PaymentMessageListener
}

func (paymentConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (paymentConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h paymentConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: topic=%s partition=%d offset=%d key=%s value=%s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		sess.MarkMessage(msg, "ok")
	}

	return nil
}

func (p *PaymentResponseKafka) StartConsume() {
	for {
		err := p.consumerGroup.Consume(context.Background(), []string{"debezium.order.payment_outbox"}, paymentConsumerGroupHandler{})
		if err != nil {
			log.Panicf("error from consumer: %v", err)
		}
	}
}
