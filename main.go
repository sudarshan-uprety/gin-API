package main

import (
	"API/database"
	"API/initializers"
	routers "API/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRoutes(router)
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
