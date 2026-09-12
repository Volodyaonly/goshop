package main

import (
	"context"
	"encoding/json"
	"log"

	"goshop/internal/config"
	"goshop/internal/events"

	"github.com/segmentio/kafka-go"
)

func main() {
	cfg := config.Load()
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			cfg.KafkaBroker,
		},
		Topic:   cfg.KafkaTopic,
		GroupID: "analytics",
	})

	defer reader.Close()

	log.Println("AnalyticsService запущен")

	ctx := context.Background()

	for {
		message, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Println("ошибка чтения Kafka:", err)
			continue
		}

		var event events.OrderCreated

		if err := json.Unmarshal(message.Value, &event); err != nil {
			log.Println("невалидное событие:", err)
			continue
		}

		log.Printf(
			"Analytics: заказ #%d, товар=%d, количество=%d, сумма=%.2f",
			event.OrderID,
			event.ProductID,
			event.Quantity,
			event.Total,
		)

		if err := reader.CommitMessages(
			ctx,
			message,
		); err != nil {
			log.Println("ошибка commit:", err)
		}
	}
}
