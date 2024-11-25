package routes

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:user_id", controller.GetUser) // Assuming this gets a single user
	router.GET("/users/signup", controller.SignUp)
	router.GET("/users/login", controller.Login)
}