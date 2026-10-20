package routers

import (
	"blog-service/controllers"
	"blog-service/repositories"
	"blog-service/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRouter struct {
	Api *gin.RouterGroup
	DB  *pgxpool.Pool
}

func NewUserRouter(api *gin.RouterGroup, db *pgxpool.Pool) *UserRouter {
	return &UserRouter{
		Api: api,
		DB:  db,
	}
}

func (r *UserRouter) Init() {
	userRepo := repositories.NewUserRepository(r.DB)
	userService := services.NewUserService(r.DB, userRepo)
	userController := controllers.NewUserController(userRepo, userService)

	userRouter := r.Api.Group("/users")
	{
		userRouter.POST("", userController.CreateUser)
		userRouter.GET("", userController.GetAllUsers)
		userRouter.GET("/:id", userController.GetUserByID)
	}
}
