package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/swapnil/restaurant-management/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
incomingRoutes.GET("/orderItems-order/:orderItem_id", controllers.GetOrderItemsByOrder())
incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}