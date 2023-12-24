package routers

import (
	controller "API/controllers"
	"API/middleware"

	"github.com/gin-gonic/gin"
)

func PostRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/posts", middleware.RequiredAuth(), controller.GetPosts())
	incomingRoutes.POST("/post", middleware.RequiredAuth(), controller.CreatePost())
	incomingRoutes.PUT("/post/:id", middleware.RequiredAuth(), controller.UpdatePost())
	incomingRoutes.DELETE("/post/:id", middleware.RequiredAuth(), controller.DeletePost())
}
