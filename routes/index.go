package routes

import (
	"ChatApp/routes/users"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func userRoutes(router *gin.RouterGroup, db *mongo.Client) {
	users.SetupUserRoutes(router, db)

}

func healthCheckRoute(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}

func noPageFound(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "Route not found",
	})
	log.Printf("%s is not found", c.Request.URL.Path)
}

func SetupRouter(engine *gin.Engine, db *mongo.Client) {
	apiGroup := engine.Group("/api")
	userRoutes(apiGroup, db)
	healthCheckRoute(engine.Group("/"))
	engine.NoRoute(noPageFound)

}
