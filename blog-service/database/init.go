package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitTables(ctx context.Context, db *pgxpool.Pool) error {
	statements := []string{
		`CREATE EXTENSION IF NOT EXISTS pgcrypto`,

		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			full_name VARCHAR(255) NOT NULL,
			user_name VARCHAR(100) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			date_of_birth DATE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,

		`CREATE TABLE IF NOT EXISTS categories (
			id UUID PRIMARY KEY,
			name VARCHAR(150) NOT NULL,
			path VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,

		`CREATE TABLE IF NOT EXISTS blogs (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			author_id UUID NOT NULL REFERENCES users(id)
				ON DELETE CASCADE,
			img_url TEXT,
			category_id UUID REFERENCES categories(id)
				ON DELETE SET NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,

		`CREATE TABLE IF NOT EXISTS projects (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			url TEXT NOT NULL,
			image TEXT,
			category_id UUID REFERENCES categories(id)
				ON DELETE SET NULL,
			user_id UUID NOT NULL REFERENCES users(id)
				ON DELETE CASCADE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)`,

		`CREATE INDEX IF NOT EXISTS idx_blogs_author_id
			ON blogs(author_id)`,

		`CREATE INDEX IF NOT EXISTS idx_blogs_category_id
			ON blogs(category_id)`,

		`CREATE INDEX IF NOT EXISTS idx_projects_category_id
			ON projects(category_id)`,

		`CREATE INDEX IF NOT EXISTS idx_projects_user_id
			ON projects(user_id)`,
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin init tables: %w", err)
	}
	defer tx.Rollback(ctx)

	for _, statement := range statements {
		if _, err := tx.Exec(ctx, statement); err != nil {
			return fmt.Errorf("execute init SQL: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit init tables: %w", err)
	}

	return nil
}
