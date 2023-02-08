package routes

import (
	controller_state "tutorial/controllers/state"

	"github.com/gin-gonic/gin"
)

func StateRoute(router *gin.RouterGroup) {
	router.GET("/state/:nebulizer_id", controller_state.GetState())
	router.PUT("/state/:nebulizer_id", controller_state.UpdateState())
}
