package events

const UserCreatedType = "user.created"
const UserCreatedQueue = "blog-sale.queue"

type UserCreated struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
