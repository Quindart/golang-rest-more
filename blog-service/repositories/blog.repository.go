package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogRepository struct {
	DB *pgxpool.Pool
}

type BlogResponse struct {
	Id           string    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	ImgUrl       string    `json:"img_url"`
	CreatedAt    time.Time `json:"created_at"`
	AuthorName   string    `json:"author_name"`
	CategoryName string    `json:"category_name"`
	CategoryPath string    `json:"category_path"`
}

func NewBlogRepository(db *pgxpool.Pool) *BlogRepository {
	return &BlogRepository{
		DB: db,
	}
}

func (repo *BlogRepository) GetAll(
	ctx context.Context,
) ([]BlogResponse, error) {
	query := `
        SELECT
            b.title,
            b.created_at,
            b.id,
            b.img_url,
            b."content",
            COALESCE(u.full_name, ''),
            COALESCE(c."path", ''),
            COALESCE(c."name", '')
        FROM blogs b
        LEFT JOIN users u ON b.author_id = u.id
        LEFT JOIN categories c ON c.id = b.category_id
        ORDER BY b.created_at DESC
    `

	rows, err := repo.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query blogs: %w", err)
	}
	defer rows.Close()

	blogs := make([]BlogResponse, 0)

	for rows.Next() {
		var blog BlogResponse
		err := rows.Scan(
			&blog.Title,
			&blog.CreatedAt,
			&blog.Id,
			&blog.ImgUrl,
			&blog.Content,
			&blog.AuthorName,
			&blog.CategoryPath,
			&blog.CategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("scan blog: %w", err)
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate blogs: %w", err)
	}
	return blogs, nil
}
