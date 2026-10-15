package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	DB *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

type ProjectResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Image        string `json:"image"`
	AuthorName   string `json:"author_name"`
	CategoryName string `json:"category_name"`
}

func (repo *ProjectRepository) GetAll(
	ctx context.Context,
) ([]ProjectResponse, error) {
	query := `
		SELECT
			p.id,
			p."name",
			p.url,
			p.image,
			COALESCE(u.full_name, '') AS author_name,
			COALESCE(c."name", '') AS category_name
		FROM projects p
		LEFT JOIN users u ON u.id = p.user_id
		LEFT JOIN categories c ON c.id = p.category_id
		ORDER BY p.id DESC
	`

	rows, err := repo.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]ProjectResponse, 0)

	for rows.Next() {
		var project ProjectResponse

		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Url,
			&project.Image,
			&project.AuthorName,
			&project.CategoryName,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
