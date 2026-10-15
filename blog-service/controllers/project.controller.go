package controllers

import (
	"blog-service/repositories"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	Repo *repositories.ProjectRepository
}

func NewProjectController(repo *repositories.ProjectRepository) *ProjectController {
	return &ProjectController{
		Repo: repo,
	}
}

func (handler *ProjectController) GetAllProjects(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	projects, err := handler.Repo.GetAll(ctx)

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch blog: " + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{
		"message":  "Get all projects successfully !",
		"projects": projects,
		"status":   http.StatusOK,
	})
}
