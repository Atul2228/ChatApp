package main

import (
	"ChatApp/config"
	"ChatApp/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Print("this")
	db := config.ConnectDB(cfg.DatabaseURL)
	engine := gin.Default()

	routes.SetupRouter(engine, db)

	err := engine.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Failed to run server: %s", err)
	} else {
		log.Println("Server is running on port", cfg.Port)
	}
}
