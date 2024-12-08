package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/swapnil/restaurant-management/database"
	"github.com/swapnil/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


var foodCollection *mongo.Collection
var menuCollection *mongo.Collection

func init() {
    ctx := context.TODO() // Create a context
    foodCollection = database.OpenCollection(database.Client, ctx, "food")
    menuCollection = database.OpenCollection(database.Client, ctx, "menu")
}



var validate = validator.New()
func GetFoods() gin.HandlerFunc{

	return func(c *gin.Context){
     
	  
	}
}

func GetFood() gin.HandlerFunc{

	return func(c*gin.Context){
		var ctx, cancel =  context.WithTimeout(context.Background(), 100*time.Second)
		foodId:=c.Param("food_id")
		var food models.Food
	   err:= foodCollection.FindOne(ctx,bson.M{"food_id":foodId}).Decode(&food)
  
		defer cancel()
		if err !=nil {
		  c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching the food item"})
		}
	}
}
func CreateFood()gin.HandlerFunc{

	return func(c*gin.Context){
var ctx,cancel = context.WithTimeout(context.Background(), 100*time.Second)
var menu models.Menu
var food models.Food

if err := c.BindJSON(&food); err !=nil{
c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
}

validationErr:=validate.Struct(food)

if validationErr !=nil{
	c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
	return
}

err := menuCollection.FindOne(ctx,bson.M{"menu_id":food.Menu_id}).Decode(&menu)
defer cancel()
if err != nil{
	msg:=fmt.Sprintf("menu was not found")
	c.JSON(http.StatusInternalServerError,gin.H{"error": msg})
}
food.Created_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
food.Updated_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
food.ID = primitive.NewObjectID()
food.Food_id= food.ID.Hex()
var num = toFixed(*food.Price,2)
food.Price = &num


}
}

func round(numb float64)int {

}

func toFixed (namb float64) int{
	
}

func UpdateFood()gin.HandlerFunc{

	return func (c* gin.Context){

	}
}