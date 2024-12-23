package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID    primitive.ObjectID   `bson:"_id"`
	Name  *string    `json:"name" valideate:"required,min=2,max=100`
	Price *float64  `json:"price" validate:"required"`
	Food_image *string `json:"food_image" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"udpated_at`
	Menu_id  *string  `json:"menu_id" validate:"required`
	Food_id *string    `json:"food_id"`
}