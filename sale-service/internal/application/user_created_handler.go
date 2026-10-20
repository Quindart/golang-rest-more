package application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/quindart/sale-service/internal/event"
	"github.com/quindart/sale-service/internal/model"
	"github.com/quindart/sale-service/internal/repository"
)

type UserCreatedHandler struct {
	CustomerRepo *repository.CustomerRepository
}

func NewUserCreatedHandler(customerRepo *repository.CustomerRepository) *UserCreatedHandler {
	return &UserCreatedHandler{CustomerRepo: customerRepo}
}

func (h *UserCreatedHandler) Handle(ctx context.Context, eventID string, message event.UserCreated) error {
	userID, err := uuid.Parse(message.UserID)
	if err != nil {
		return fmt.Errorf("invalid user_id: %w", err)
	}
	if message.Email == "" || message.PhoneNumber == "" {
		return fmt.Errorf("email and phone_number are required")
	}
	return h.CustomerRepo.Create(ctx, model.Customer{UserID: userID, Email: message.Email, PhoneNumber: message.PhoneNumber}, eventID)
}
