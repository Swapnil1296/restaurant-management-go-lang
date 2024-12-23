package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/swapnil/restaurant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
incomingRoutes.GET("/invoices", controllers.GetInvoices())
incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
incomingRoutes.POST("/invoices", controllers.CreateInvoices())
incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoices())
}