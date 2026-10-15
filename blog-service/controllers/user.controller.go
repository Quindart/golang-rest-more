package controllers

import (
	"blog-service/repositories"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func NewUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{
		Repo: repo,
	}
}

func (handler *UserController) GetAllUsers(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	users, err := handler.Repo.GetAll(ctx)

	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch user: " + err.Error(),
		})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Get all users successfully !",
		"users":   users,
	})
}
