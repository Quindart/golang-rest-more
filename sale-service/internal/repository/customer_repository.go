package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quindart/sale-service/internal/model"
)

type CustomerRepository struct {
	DB *pgxpool.Pool
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
