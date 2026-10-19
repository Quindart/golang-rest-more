package mq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

const QueueName = "demo.tasks"

func Connect() (
	*amqp.Connection,
	*amqp.Channel,
	amqp.Queue,
	error,
) {
	// Đọc URL từ biến môi trường.
	url := os.Getenv("RABBITMQ_URL")

	// Nếu không cấu hình, sử dụng RabbitMQ local.
	if url == "" {
		url = "amqp://app:app_password@localhost:5672/"
	}

	// Bước 1: Tạo kết nối đến RabbitMQ.
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, amqp.Queue{},
			fmt.Errorf("connect to RabbitMQ: %w", err)
	}

	// Bước 2: Mở AMQP channel.
	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()

		return nil, nil, amqp.Queue{},
			fmt.Errorf("open channel: %w", err)
	}

	// Bước 3: Khai báo queue.
	// Producer và Consumer phải dùng cấu hình tương thích.
	queue, err := ch.QueueDeclare(
		QueueName, // Tên queue
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)

	if err != nil {
		_ = ch.Close()
		_ = conn.Close()

		return nil, nil, amqp.Queue{},
			fmt.Errorf("declare queue: %w", err)
	}

	return conn, ch, queue, nil
}
