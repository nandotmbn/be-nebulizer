package routes

import (
	controller_nebulizer "tutorial/controllers/nebulizer"

	"github.com/gin-gonic/gin"
)

func NebulizerRoute(router *gin.RouterGroup) { // http://localhost:8080/v1
	router.POST("/nebulizer", controller_nebulizer.RegisterNebulizer()) // http://localhost:8080/v1/nebulizer
	router.POST("/nebulizer/retriveid", controller_nebulizer.GetIdNebulizer())
}
