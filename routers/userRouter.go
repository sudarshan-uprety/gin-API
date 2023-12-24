package routers

import (
	controller "API/controllers"
	"API/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controller.RegisterUser())
	incomingRoutes.POST("/login", controller.Login())
	incomingRoutes.GET("/validate", middleware.RequiredAuth(), controller.Validate())
}
