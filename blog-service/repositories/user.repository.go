package repositories

import (
	"blog-service/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, tx pgx.Tx, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	user.Password = string(hashedPassword)
	_, err = tx.Exec(ctx, `
		INSERT INTO users (id, full_name, user_name, password, date_of_birth, phone_number, email)
		VALUES ($1::uuid, $2, $3, $4, NULLIF($5, '')::date, $6, $7)
	`, user.Id, user.FullName, user.UserName, user.Password, user.DateOfBirth, user.PhoneNumber, user.Email)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT
			id::text,
			full_name,
			user_name,
			COALESCE(phone_number, ''),
			COALESCE(email, ''),
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
			&user.PhoneNumber,
			&user.Email,
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
			COALESCE(phone_number, ''),
			COALESCE(email, ''),
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
		&user.PhoneNumber,
		&user.Email,
		&user.DateOfBirth,
	)
	if err != nil {
		return nil, fmt.Errorf("get user by ID: %w", err)
	}

	return &user, nil
}
