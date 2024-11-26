package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/swapnil/restaurant-management/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:user_id", controllers.GetUser) // Assuming this gets a single user
	router.GET("/users/signup", controllers.SignUp)
	router.GET("/users/login", controllers.Login)
}