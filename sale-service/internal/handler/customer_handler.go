package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quindart/sale-service/internal/client"
	"github.com/quindart/sale-service/internal/model"
	"github.com/quindart/sale-service/internal/repository"
)

type CustomerHandler struct {
	CustomerRepo *repository.CustomerRepository
	UserClient   *client.UserClient
}

func NewCustomerHandler(
	customerRepo *repository.CustomerRepository,
	userClient *client.UserClient,
) *CustomerHandler {
	return &CustomerHandler{
		CustomerRepo: customerRepo,
		UserClient:   userClient,
	}
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	ctx := c.Request().Context()
	customers, err := h.CustomerRepo.GetAllCustomers(ctx)
	if err != nil {
		c.Logger().Errorf("get customers: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get customers",
		})
	}

	if len(customers) == 0 {
		return c.JSON(http.StatusOK, []model.CustomerResponse{})
	}

	result := make([]model.CustomerResponse, 0, len(customers))
	users := make(map[string]*client.UserResponse, len(customers))

	for _, customer := range customers {
		user, exists := users[customer.UserID.String()]
		if !exists && h.UserClient != nil {
			user, err = h.UserClient.GetUserByID(ctx, customer.UserID.String())
			if err != nil {
				c.Logger().Errorf("get user %s: %v", customer.UserID, err)
				return c.JSON(http.StatusBadGateway, map[string]string{"message": "Failed to retrieve user information"})
			}
			users[customer.UserID.String()] = user
		}
		if user == nil {
			return fmt.Errorf("user profile is unavailable for customer %s", customer.ID)
		}

		result = append(result, model.CustomerResponse{
			ID:          customer.ID,
			UserID:      user.ID,
			FullName:    user.FullName,
			PhoneNumber: customer.PhoneNumber,
			Email:       customer.Email,
		})
	}

	return c.JSON(http.StatusOK, result)
}
