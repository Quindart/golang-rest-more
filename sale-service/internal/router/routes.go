package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/quindart/sale-service/internal/handler"
)

func Register(e *echo.Echo, customerHandler *handler.CustomerHandler) {
	api := e.Group("/api/v1")
	api.GET("/customers", customerHandler.GetAllCustomers)
}
