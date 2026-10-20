package controllers

import (
	"blog-service/models"
	"blog-service/repositories"
	"blog-service/services"
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserController struct {
	Repo    *repositories.UserRepository
	Service *services.UserService
}

func NewUserController(repo *repositories.UserRepository, service *services.UserService) *UserController {
	return &UserController{
		Repo: repo, Service: service,
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

type createUserRequest struct {
	FullName    string `json:"full_name" binding:"required"`
	UserName    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required,min=8"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
}

func (handler *UserController) CreateUser(c *gin.Context) {
	var request createUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data: " + err.Error()})
		return
	}
	request.FullName = strings.TrimSpace(request.FullName)
	request.UserName = strings.TrimSpace(request.UserName)
	request.PhoneNumber = strings.TrimSpace(request.PhoneNumber)
	request.Email = strings.TrimSpace(request.Email)
	if request.FullName == "" || request.UserName == "" || request.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "full_name, user_name, and phone_number are required"})
		return
	}
	if request.DateOfBirth != "" {
		if _, err := time.Parse("2006-01-02", request.DateOfBirth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "date_of_birth must use YYYY-MM-DD format"})
			return
		}
	}

	user := models.NewUser(request.FullName, request.UserName, request.Password,
		request.DateOfBirth, request.PhoneNumber, request.Email)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	if err := handler.Service.Create(ctx, user); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{"message": "user_name or email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (handler *UserController) GetUserByID(ginContext *gin.Context) {
	id := ginContext.Param("id")

	ctx, cancel := context.WithTimeout(
		ginContext.Request.Context(),
		5*time.Second,
	)
	defer cancel()

	user, err := handler.Repo.GetByID(ctx, id)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch user",
		})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Get user successfully",
		"user":    user,
	})
}
