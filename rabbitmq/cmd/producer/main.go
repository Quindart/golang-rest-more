package main

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"rabbitmq-go-demo/internal/mq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, ch, queue, err := mq.Connect()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer ch.Close()

	// Lấy danh sách message từ command line.
	messages := os.Args[1:]

	// Nếu không truyền message, gửi một message mặc định.
	if len(messages) == 0 {
		messages = []string{
			"Hello RabbitMQ from Go",
		}
	}

	for _, message := range messages {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)

		err := ch.PublishWithContext(
			ctx,
			"",         // Exchange mặc định
			queue.Name, // Routing key
			false,      // Mandatory
			false,      // Immediate
			amqp.Publishing{
				ContentType:  "text/plain",
				DeliveryMode: amqp.Persistent,
				Timestamp:    time.Now(),
				Body:         []byte(message),
			},
		)

		cancel()

		if err != nil {
			return err
		}

		log.Printf("[producer] published: %q", message)
	}

	return nil
}
