package routers

import (
	"blog-service/controllers"
	"blog-service/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogRoute struct {
	Api *gin.RouterGroup
	DB  *pgxpool.Pool
}

func NewBlogRouter(api *gin.RouterGroup, db *pgxpool.Pool) *BlogRoute {
	return &BlogRoute{
		Api: api,
		DB:  db,
	}
}

func (r *BlogRoute) Init() {
	blogRepo := repositories.NewBlogRepository(r.DB)
	blogController := controllers.NewBlogController(blogRepo)

	blogRouter := r.Api.Group("/blogs")
	{
		blogRouter.GET("", blogController.GetAllBlogs)
	}
}
