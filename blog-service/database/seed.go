package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func Seed(ctx context.Context, db *pgxpool.Pool) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin seed: %w", err)
	}
	defer tx.Rollback(ctx)

	userID := uuid.New()
	categoryID := uuid.New()
	projectCategoryID := uuid.New()
	blogID := uuid.New()
	projectID := uuid.New()

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte("ChangeMe123!"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("hash seed password: %w", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO users
			(id, full_name, user_name, password, date_of_birth)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_name) DO NOTHING
	`, userID, "Demo User", "demo_user", string(passwordHash), "2000-01-15")
	if err != nil {
		return fmt.Errorf("seed user: %w", err)
	}

	// Retrieve the actual user ID if this username already exists.
	err = tx.QueryRow(ctx, `
		SELECT id FROM users WHERE user_name = $1
	`, "demo_user").Scan(&userID)
	if err != nil {
		return fmt.Errorf("get seed user: %w", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO categories (id, name, path)
		VALUES
			($1, 'Technology', 'technology'),
			($2, 'Backend', 'backend')
		ON CONFLICT (path) DO NOTHING
	`, categoryID, projectCategoryID)
	if err != nil {
		return fmt.Errorf("seed categories: %w", err)
	}

	err = tx.QueryRow(ctx, `
		SELECT id FROM categories WHERE path = $1
	`, "technology").Scan(&categoryID)
	if err != nil {
		return fmt.Errorf("get technology category: %w", err)
	}

	err = tx.QueryRow(ctx, `
		SELECT id FROM categories WHERE path = $1
	`, "backend").Scan(&projectCategoryID)
	if err != nil {
		return fmt.Errorf("get backend category: %w", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO blogs
			(id, title, content, author_id, img_url, category_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO NOTHING
	`, blogID, "Getting Started with Go",
		"An introduction to building backend applications with Go.",
		userID, "https://example.com/go.jpg", categoryID)
	if err != nil {
		return fmt.Errorf("seed blog: %w", err)
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO projects
			(id, name, url, image, category_id, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO NOTHING
	`, projectID, "Go MVC Blog",
		"https://github.com/example/go-mvc-blog",
		"https://example.com/project.jpg", projectCategoryID, userID)
	if err != nil {
		return fmt.Errorf("seed project: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit seed: %w", err)
	}

	return nil
}
