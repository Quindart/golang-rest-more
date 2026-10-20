package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quindart/sale-service/internal/client"
	"github.com/quindart/sale-service/internal/database"
	"github.com/quindart/sale-service/internal/handler"
	"github.com/quindart/sale-service/internal/repository"
	routes "github.com/quindart/sale-service/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		log.Fatal("USER_SERVICE_URL is required")
	}

	dbCtx, cancelDB := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	db, err := database.NewPostgresPool(dbCtx)
	cancelDB()

	if err != nil {
		log.Fatalf("initialize database: %v", err)
	}
	defer db.Close()

	customerRepo := repository.NewCustomerRepository(db)
	userClient := client.NewUserClient(userServiceURL)

	customerHandler := handler.NewCustomerHandler(
		customerRepo,
		userClient,
	)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Allow all origins
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	routes.Register(e, customerHandler)
	
	go func() {
		if err := e.Start(":8083"); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			e.Logger.Errorf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Errorf("shutdown error: %v", err)
	}
}
