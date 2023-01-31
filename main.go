package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanoviedog/golang-ms/controllers"
	"github.com/juanoviedog/golang-ms/initializers"
	"github.com/juanoviedog/golang-ms/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run()
}
