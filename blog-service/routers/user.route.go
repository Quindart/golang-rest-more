package routers

import (
	"blog-service/controllers"
	"blog-service/repositories"

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
	userController := controllers.NewUserController(userRepo)

	userRouter := r.Api.Group("/users")
	{
		userRouter.GET("", userController.GetAllUsers)
		userRouter.GET("/:id", userController.GetUserByID)
	}
}
