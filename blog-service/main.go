package main

import (
	"blog-service/database"
	"blog-service/infra"
	"blog-service/infra/rabbitmq"
	"blog-service/routers"
	"blog-service/services"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env: ", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	pool, err := infra.NewPostgres(ctx, databaseURL)

	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	log.Println("PostgreSQL connected successfully")
	// Initialize tables
	if err := database.InitTables(ctx, pool); err != nil {
		log.Fatal("Failed to initialize tables: ", err)
	}
	log.Println("Database tables initialized successfully")

	// Seed sample data
	if err := database.Seed(ctx, pool); err != nil {
		log.Fatal("Failed to seed database: ", err)
	}
	log.Println("Database seeded successfully")

	conn, ch, err := rabbitmq.Connect("")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	defer ch.Close()

	appCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go (&services.OutboxService{DB: pool, Publisher: rabbitmq.NewPublisher(ch)}).Run(appCtx)

	r := routers.SetupRouter(pool)
	if err := r.Run(":8082"); err != nil {
		log.Fatal("HTTP server failed: ", err)
	}

}
