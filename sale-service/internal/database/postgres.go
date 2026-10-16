package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context) (*pgxpool.Pool, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database config: %w", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := initTable(ctx, pool); err != nil {
		pool.Close()
		return nil, fmt.Errorf("initialize tables: %w", err)
	}

	return pool, nil
}

func initTable(ctx context.Context, db *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS customers (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			phone_number VARCHAR(20),
			user_id UUID,
			email VARCHAR(255)
		);
	`
	if _, err := db.Exec(ctx, query); err != nil {
		return fmt.Errorf("create customers table: %w", err)
	}
	return nil
}
