package routes

import (
	"github.com/erkylima/allGolangRithm/api/v1/algorithms/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesRegistry(r *gin.Engine) {
	r.POST("/add_two_numbers", handlers.AddTwoNumbers)
}
