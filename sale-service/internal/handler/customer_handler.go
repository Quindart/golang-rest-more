package handler

import (
	"net/http"

	"github.com/google/uuid"
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

	// 1. Get customers from Customer Service database.
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

	// 2. Collect unique user UUIDs.
	uniqueIDs := make(map[string]struct{})
	userIDs := make([]uuid.UUID, 0, len(customers))

	for _, customer := range customers {
		idKey := customer.UserID.String()

		if _, exists := uniqueIDs[idKey]; exists {
			continue
		}

		uniqueIDs[idKey] = struct{}{}
		userIDs = append(userIDs, customer.UserID)
	}

	// 3. Call User Service once.
	result := make([]model.CustomerResponse, 0, len(customers))

	for _, customer := range customers {
		user, err := h.UserClient.GetUserByID(
			ctx,
			customer.UserID.String(),
		)
		if err != nil {
			c.Logger().Errorf(
				"get user %s: %v",
				customer.UserID,
				err,
			)

			return c.JSON(http.StatusBadGateway, map[string]string{
				"message": "Failed to retrieve user information",
			})
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
