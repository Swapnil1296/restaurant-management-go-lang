package main

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
	"github.com/swapnil/restaurant-management/routes"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

port:=os.Getenv("PORT")
// If the port is not set in the environment, we'll use 8080 as the default port
if port==""{
	port="8080"
}
// Create a new router using the Gin framework
router := gin.New()
router.Use(gin.Logger())

// Define the route for the index page and display the Go version
routes.UserRoutes(router)
middleware.AuthMiddleware(router)
router.Use(middleware.AuthMiddleware())
routes.foodRouter(router)



// Start serving the application
router.Run(":" + port)
fmt.Println("Server is running on port: " + port)

}