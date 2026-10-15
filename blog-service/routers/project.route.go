package routers

import (
	"blog-service/controllers"
	"blog-service/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRouter struct {
	Api *gin.RouterGroup
	DB  *pgxpool.Pool
}

func NewProjectRouter(api *gin.RouterGroup, db *pgxpool.Pool) *ProjectRouter {
	return &ProjectRouter{
		Api: api,
		DB:  db,
	}
}

func (r *ProjectRouter) Init() {
	projectRepo := repositories.NewProjectRepository(r.DB)
	projectController := controllers.NewProjectController(projectRepo)

	projectRouter := r.Api.Group("/projects")
	{
		projectRouter.GET("", projectController.GetAllProjects)
	}
}
