package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"goshop/internal/events"
	"goshop/internal/models"

	kafkago "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafkago.Writer
}

func NewProducer(broker string, topic string) *Producer {
	writer := &kafkago.Writer{
		Addr:     kafkago.TCP(broker),
		Topic:    topic,
		Balancer: &kafkago.LeastBytes{},
	}

	return &Producer{
		writer: writer,
	}
}

func (p *Producer) PublishOrderCreated(
	ctx context.Context,
	order *models.Order,
) error {

	event := events.OrderCreated{
		EventID:   fmt.Sprintf("order_created:%d", order.ID),
		Event:     "order_created",
		OrderID:   order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Total:     order.Total,
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(
		ctx,
		kafkago.Message{
			Key:   []byte(strconv.Itoa(order.ID)),
			Value: data,
		},
	)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
