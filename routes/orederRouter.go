package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/swapnil/restaurant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
incomingRoutes.GET("/orders", controllers.GetOrders())
incomingRoutes.GET("/orders/:order_id", controllers.GetOrder())
incomingRoutes.POST("/orders", controllers.CreateMenu())
incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateMenu())
}