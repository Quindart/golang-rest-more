package main

import (
	"blog-service/database"
	"blog-service/infra"
	"blog-service/routers"
	"context"
	"log"
	"os"
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
	r := routers.SetupRouter(pool)
	r.Run(":8082")

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

}
