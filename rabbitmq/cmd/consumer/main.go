package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Chỉ cho phép một message chưa ACK
	// tại một thời điểm trên channel này.
	if err := ch.Qos(1, 0, false); err != nil {
		return fmt.Errorf("set QoS: %w", err)
	}

	consumerTag := fmt.Sprintf("worker-%d", os.Getpid())

	// Đăng ký Consumer với queue.
	messages, err := ch.Consume(
		queue.Name,
		consumerTag,
		false, // Manual ACK
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return fmt.Errorf("register consumer: %w", err)
	}

	// Bắt tín hiệu dừng chương trình.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(stop)

	log.Printf(
		"[consumer] waiting on queue %q; press Ctrl+C to stop",
		queue.Name,
	)

	for {
		select {
		case <-stop:
			log.Println("[consumer] stopping")

			if err := ch.Cancel(consumerTag, false); err != nil {
				log.Printf("cancel consumer: %v", err)
			}

			return nil

		case delivery, ok := <-messages:
			if !ok {
				return fmt.Errorf("consumer channel closed")
			}

			log.Printf(
				"[consumer] received: %q, redelivered=%t",
				string(delivery.Body),
				delivery.Redelivered,
			)

			// Giả lập xử lý công việc.
			if err := processMessage(delivery.Body); err != nil {
				log.Printf("process message: %v", err)
				continue
			}

			// Xác nhận xử lý thành công.
			if err := delivery.Ack(false); err != nil {
				log.Printf("ACK failed: %v", err)
				continue
			}

			log.Printf(
				"[consumer] processed and ACKed: %q",
				string(delivery.Body),
			)
		}
	}
}

func processMessage(body []byte) error {
	log.Printf("[worker] processing: %s", string(body))

	// Giả lập công việc mất một giây.
	time.Sleep(time.Second)

	log.Printf("[worker] finished: %s", string(body))

	return nil
}