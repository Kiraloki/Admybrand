// main.go
package main

import (
	database "version1/database"
	handlers "version1/hanlders"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the MongoDB connection
	database.InitMongoDB()

	// Routes
	r.POST("/models", handlers.CreateModel)
	r.GET("/models", handlers.GetModels)
	r.GET("/models/:id", handlers.GetModelByID)
	r.PUT("/models/:id", handlers.UpdateModel)
	r.DELETE("/models/:id", handlers.DeleteModel)


	// Run the server
	r.Run(":8080")
}
