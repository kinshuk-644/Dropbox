package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/kinshuk-644/Dropbox/backend/config"
	"github.com/kinshuk-644/Dropbox/backend/connection"
	"github.com/kinshuk-644/Dropbox/backend/handlers"
)

func main() {
	config.LoadConfig()

	connection.ConnectToDB()

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	router.Use(cors.New(corsConfig))

	router.POST("/upload", handlers.UploadFile)
	router.GET("/files", handlers.GetFiles)
	router.GET("/files/:id", handlers.DownloadFile)

	port := "8000"
	fmt.Printf("✅ Backend server starting on http://localhost:%s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
