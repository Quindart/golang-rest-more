package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quindart/sale-service/internal/model"
)

type CustomerRepository struct {
	DB *pgxpool.Pool
}

func (r *CustomerRepository) Create(ctx context.Context, customer model.Customer, eventID string) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin create customer: %w", err)
	}
	defer tx.Rollback(ctx)

	var inserted bool
	err = tx.QueryRow(ctx, `
		INSERT INTO consumed_events (event_id, event_type)
		VALUES ($1::uuid, 'user.created')
		ON CONFLICT (event_id) DO NOTHING
		RETURNING true
	`, eventID).Scan(&inserted)
	if err != nil && err != pgx.ErrNoRows {
		return fmt.Errorf("record consumed event: %w", err)
	}
	if err == nil {
		_, err = tx.Exec(ctx, `
			INSERT INTO customers (user_id, phone_number, email)
			VALUES ($1, $2, $3)
			ON CONFLICT (user_id) DO UPDATE
			SET phone_number = EXCLUDED.phone_number, email = EXCLUDED.email
		`, customer.UserID, customer.PhoneNumber, customer.Email)
		if err != nil {
			return fmt.Errorf("insert customer: %w", err)
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit create customer: %w", err)
	}
	return nil
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{
		DB: db,
	}
}

func (r *CustomerRepository) GetAllCustomers(
	ctx context.Context,
) ([]model.Customer, error) {
	query := `
		SELECT id, phone_number, user_id, email
		FROM customers
		ORDER BY id ASC
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query customers: %w", err)
	}
	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {
		var customer model.Customer

		if err := rows.Scan(
			&customer.ID,
			&customer.PhoneNumber,
			&customer.UserID,
			&customer.Email,
		); err != nil {
			return nil, fmt.Errorf("scan customer: %w", err)
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate customers: %w", err)
	}

	return customers, nil
}
