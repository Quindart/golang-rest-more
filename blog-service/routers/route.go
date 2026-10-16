package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
	}))

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
