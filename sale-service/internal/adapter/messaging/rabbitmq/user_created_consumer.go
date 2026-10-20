package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/quindart/sale-service/internal/application"
	"github.com/quindart/sale-service/internal/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

type UserCreatedConsumer struct {
	RabbitURL string
	Handler   *application.UserCreatedHandler
}

func NewUserCreatedConsumer(userHandler *application.UserCreatedHandler) *UserCreatedConsumer {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://app:app_password@localhost:5672/"
	}
	return &UserCreatedConsumer{RabbitURL: url, Handler: userHandler}
}

func (c *UserCreatedConsumer) Run(ctx context.Context) {
	for ctx.Err() == nil {
		if err := c.consume(ctx); err != nil && ctx.Err() == nil {
			log.Printf("user-created consumer: %v", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
		}
	}
}

func (c *UserCreatedConsumer) consume(ctx context.Context) error {
	conn, err := amqp.Dial(c.RabbitURL)
	if err != nil {
		return fmt.Errorf("connect RabbitMQ: %w", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("open channel: %w", err)
	}
	defer ch.Close()
	if _, err := ch.QueueDeclare(event.UserCreatedQueue, true, false, false, false, nil); err != nil {
		return fmt.Errorf("declare queue: %w", err)
	}
	if err := ch.Qos(1, 0, false); err != nil {
		return fmt.Errorf("set prefetch: %w", err)
	}
	deliveries, err := ch.Consume(event.UserCreatedQueue, "sale-service", false, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("consume queue: %w", err)
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case delivery, ok := <-deliveries:
			if !ok {
				return fmt.Errorf("delivery channel closed")
			}
			if err := c.handle(ctx, delivery); err != nil {
				log.Printf("handle event %s: %v", delivery.MessageId, err)
				if nackErr := delivery.Nack(false, false); nackErr != nil {
					return nackErr
				}
				continue
			}
			if err := delivery.Ack(false); err != nil {
				return fmt.Errorf("ack message: %w", err)
			}
		}
	}
}

func (c *UserCreatedConsumer) handle(ctx context.Context, delivery amqp.Delivery) error {
	if delivery.Type != event.UserCreatedType {
		return fmt.Errorf("unsupported event type %q", delivery.Type)
	}
	var message event.UserCreated
	if err := json.Unmarshal(delivery.Body, &message); err != nil {
		return fmt.Errorf("decode event: %w", err)
	}
	if delivery.MessageId == "" {
		return fmt.Errorf("message ID is required")
	}
	return c.Handler.Handle(ctx, delivery.MessageId, message)
}
