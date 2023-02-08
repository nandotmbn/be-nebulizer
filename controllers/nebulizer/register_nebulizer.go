package controller_nebulizer

import (
	"context"
	"net/http"
	"time"
	"tutorial/configs"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

var nebulizerCollection *mongo.Collection = configs.GetCollection(configs.DB, "nebulizer")

func RegisterNebulizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var nebulizer models.Nebulizer
		defer cancel()
		c.BindJSON(&nebulizer)

		if validationErr := validate.Struct(&nebulizer); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		count, err_ := nebulizerCollection.CountDocuments(ctx, bson.M{"nebulizer_name": nebulizer.NebulizerName})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count >= 1 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Meter name has been taken"})
			return
		}

		bytes, errors := bcrypt.GenerateFromPassword([]byte(nebulizer.Password), 14)
		if errors != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Password tidak valid"})
		}

		newNebulizer := models.Nebulizer{
			NebulizerName: nebulizer.NebulizerName,
			State:         false,
			Battery:       0,
			Password:      string(bytes),
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		}

		result, err := nebulizerCollection.InsertOne(ctx, newNebulizer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.NebulizerView{
			NebulizerId:   result.InsertedID,
			NebulizerName: nebulizer.NebulizerName,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})
	}
}
