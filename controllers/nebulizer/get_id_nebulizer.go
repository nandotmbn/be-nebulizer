package controller_nebulizer

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetIdNebulizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nebulizerPayload views.PayloadRetriveId
		defer cancel()
		c.BindJSON(&nebulizerPayload)

		if validationErr := validate.Struct(&nebulizerPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		var resultMeter models.Nebulizer
		var finalPayload views.FinalRetriveId
		result := nebulizerCollection.FindOne(ctx, bson.M{"nebulizer_name": nebulizerPayload.NebulizerName})
		result.Decode(&resultMeter)
		result.Decode(&finalPayload)
		err := bcrypt.CompareHashAndPassword([]byte(resultMeter.Password), []byte(nebulizerPayload.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
				"data":    "Nebulizer name or password is not valid",
			})
			return
		}

		c.JSON(http.StatusOK,
			bson.M{
				"status":  http.StatusOK,
				"message": "Success",
				"data":    finalPayload,
			},
		)
	}
}
