package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// health check
	apiV1 := router.Group("/api/v1")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	userRoute := NewUserRouter(apiV1, db)
	userRoute.Init()

	blogRoute := NewBlogRouter(apiV1, db)
	blogRoute.Init()

	projectRoute := NewProjectRouter(apiV1, db)
	projectRoute.Init()
	return router

}
