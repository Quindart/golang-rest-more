package controllers

import (
	"blog-service/repositories"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Repo *repositories.BlogRepository
}

func NewBlogController(repo *repositories.BlogRepository) *BlogController {
	return &BlogController{
		Repo: repo,
	}
}

func (handler *BlogController) GetAllBlogs(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	blogs, err := handler.Repo.GetAll(ctx)

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch blog: " + err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Get all blogs successfully !",
		"blogs":   blogs,
		"status":  http.StatusOK,
	})
}
