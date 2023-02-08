package controller_state

import (
	"context"
	"net/http"
	"time"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetState() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var nebulizerId = c.Param("nebulizer_id")

		nebulizerIdObj, err := primitive.ObjectIDFromHex(nebulizerId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		var state views.LastState
		results := nebulizerCollection.FindOne(ctx, bson.M{"_id": nebulizerIdObj})
		results.Decode(&state)

		c.JSON(http.StatusOK, bson.M{
			"status":  http.StatusOK,
			"message": "success",
			"data":    state,
		})

	}
}
