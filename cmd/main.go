package main

import (
	"text-labeler-backend/pkg/rest/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()

	r.Run("0.0.0.0:8080")
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors.Default())

	handler := handlers.NewHandler()

	router.POST("/upload", handler.PostCSV)

	return router
}
