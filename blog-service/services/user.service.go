package services

import (
	"context"
	"encoding/json"
	"fmt"

	"blog-service/events"
	"blog-service/models"
	"blog-service/repositories"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	DB   *pgxpool.Pool
	Repo *repositories.UserRepository
}

func NewUserService(db *pgxpool.Pool, repo *repositories.UserRepository) *UserService {
	return &UserService{DB: db, Repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	tx, err := s.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin create user: %w", err)
	}
	defer tx.Rollback(ctx)

	if err := s.Repo.Create(ctx, tx, user); err != nil {
		return err
	}
	eventID := uuid.NewString()
	payload, err := json.Marshal(events.UserCreated{
		UserID: user.Id, PhoneNumber: user.PhoneNumber, Email: user.Email,
	})
	if err != nil {
		return fmt.Errorf("encode user-created event: %w", err)
	}
	if _, err := tx.Exec(ctx, `
		INSERT INTO outbox_events (id, event_type, payload)
		VALUES ($1::uuid, $2, $3::jsonb)
	`, eventID, events.UserCreatedType, payload); err != nil {
		return fmt.Errorf("save user-created event: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit create user: %w", err)
	}
	return nil
}
