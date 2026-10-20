package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel *amqp.Channel
}

func NewPublisher(channel *amqp.Channel) *Publisher {
	return &Publisher{Channel: channel}
}

func (p *Publisher) Publish(ctx context.Context, eventID, eventType string, payload []byte) error {
	confirmations := p.Channel.NotifyPublish(make(chan amqp.Confirmation, 1))
	if err := p.Channel.PublishWithContext(ctx, "", "blog-sale.queue", true, false, amqp.Publishing{
		ContentType: "application/json", DeliveryMode: amqp.Persistent,
		MessageId: eventID, Type: eventType, Body: payload,
	}); err != nil {
		return fmt.Errorf("publish %s: %w", eventType, err)
	}
	select {
	case confirmation := <-confirmations:
		if !confirmation.Ack {
			return fmt.Errorf("broker rejected %s", eventType)
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}
