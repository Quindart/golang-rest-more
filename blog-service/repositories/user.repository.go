package repositories

import (
	"blog-service/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT
			id::text,
			full_name,
			user_name,
			COALESCE(date_of_birth::text, '')
		FROM users
		ORDER BY full_name;
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()
	users := make([]models.User, 0)

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.UserName,
			&user.DateOfBirth,
		)
		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	id string,
) (*models.User, error) {
	query := `
		SELECT
			id::text,
			full_name,
			user_name,
			COALESCE(date_of_birth::text, '')
		FROM users
		WHERE id = $1::uuid
		LIMIT 1;
	`

	var user models.User

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&user.Id,
		&user.FullName,
		&user.UserName,
		&user.DateOfBirth,
	)
	if err != nil {
		return nil, fmt.Errorf("get user by ID: %w", err)
	}

	return &user, nil
}
