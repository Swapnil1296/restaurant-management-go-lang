package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/swapnil/restaurant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
incomingRoutes.GET("/menus", controllers.GetMenu)
incomingRoutes.GET("/menus/:menu_id", controllers.GetMenu)
incomingRoutes.POST("/menus", controllers.CreateMenu)
incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu)
}