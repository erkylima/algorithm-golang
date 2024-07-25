package main

import (
	"log"

	"github.com/erkylima/allGolangRithm/api/v1/algorithms/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RoutesRegistry(r)
	port := "3000"
	if err := r.Run(":" + port); err != nil {
		log.Fatal("error on start server:", err)
	}
}
