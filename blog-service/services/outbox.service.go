package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventPublisher interface {
	Publish(ctx context.Context, eventID, eventType string, payload []byte) error
}

type OutboxService struct {
	DB        *pgxpool.Pool
	Publisher EventPublisher
}

type outboxEvent struct {
	ID      string
	Type    string
	Payload json.RawMessage
}

func (s *OutboxService) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if err := s.publishPending(ctx); err != nil && ctx.Err() == nil {
			fmt.Printf("publish pending outbox events: %v\n", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func (s *OutboxService) publishPending(ctx context.Context) error {
	rows, err := s.DB.Query(ctx, `SELECT id::text, event_type, payload FROM outbox_events WHERE published_at IS NULL ORDER BY created_at LIMIT 50`)
	if err != nil {
		return fmt.Errorf("query outbox: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var event outboxEvent
		if err := rows.Scan(&event.ID, &event.Type, &event.Payload); err != nil {
			return fmt.Errorf("scan outbox: %w", err)
		}
		if err := s.Publisher.Publish(ctx, event.ID, event.Type, event.Payload); err != nil {
			return err
		}
		if _, err := s.DB.Exec(ctx, `UPDATE outbox_events SET published_at = NOW() WHERE id = $1::uuid`, event.ID); err != nil {
			return fmt.Errorf("mark outbox published: %w", err)
		}
	}
	return rows.Err()
}
