package main

import (
	"ChatApp/config"
	"ChatApp/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	engine := gin.Default()

	routes.SetupRouter(engine)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Failed to run server: %s", err)
	} else {
		log.Println("Server is running on port", cfg.Port)
	}
}
