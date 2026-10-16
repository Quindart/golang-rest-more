package model

import "github.com/google/uuid"

type Customer struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
}
type CustomerResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      string `json:"userId"`
	FullName    string    `json:"fullName"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
}
