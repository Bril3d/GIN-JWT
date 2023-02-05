package main

import (
	"github.com/bril3d/gin-jwt/controllers"
	"github.com/bril3d/gin-jwt/intializers"
	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDb()
	intializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()
}
