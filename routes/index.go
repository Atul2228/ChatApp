package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "List of users",
		})
	})

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

func SetupRouter(engine *gin.Engine) {
	userRoutes(engine.Group("/api"))
	healthCheckRoute(engine.Group("/"))
	engine.NoRoute(noPageFound)

}
