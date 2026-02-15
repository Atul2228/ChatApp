package users

import (
	"ChatApp/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUserRoutes(router *gin.RouterGroup, db *mongo.Client) {
	userCtrl := &controllers.UserController{DB: db}

	router.GET("/user", userCtrl.GetUser)
	router.POST("/user", userCtrl.CreateUser)
}
