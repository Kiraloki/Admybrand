
package handlers

import (
	"context"
	"log"
	"net/http"
	"time"
	database "version1/database"
	models "version1/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateModel handles the creation of a new model
func CreateModel(c *gin.Context) {
	var model models.Model
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	model.ID = primitive.NewObjectID().Hex()
	model.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := database.GetMongoDBClient() 
	collection := client.Database("Admybrand").Collection("Users")  




	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		log.Println("Failed to create the model:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the model"})
		return
	}

	c.JSON(http.StatusCreated, model)
}

// GetModels handles fetching all models
func GetModels(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := database.GetMongoDBClient()
	collection := client.Database("Admybrand").Collection("Users")
	

	cursor, err := collection.Find(ctx, bson.M{})	
	if err != nil {
		log.Println("Failed to fetch models:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch models"})
		return
	}
	defer cursor.Close(ctx)

	var models []models.Model
	if err := cursor.All(ctx, &models); err != nil {
		log.Println("Failed to decode models:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode models"})
		return
	}

	c.JSON(http.StatusOK, models)
}

// GetModelByID handles fetching a model by ID
func GetModelByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := database.GetMongoDBClient()
	collection := client.Database("Admybrand").Collection("Users")

	var model models.Model
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
			return
		}
		log.Println("Failed to fetch model:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch model"})
		return
	}

	c.JSON(http.StatusOK, model)
}

// UpdateModel handles updating a model
func UpdateModel(c *gin.Context) {
	id := c.Param("id")

	var model models.Model
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := database.GetMongoDBClient()
	collection := client.Database("Admybrand").Collection("Users")

	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": model})
	if err != nil {
		log.Println("Failed to update the model:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the model"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model updated successfully"})
}

// DeleteModel handles deleting a model
func DeleteModel(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := database.GetMongoDBClient()
	collection := client.Database("Admybrand").Collection("Users")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("Failed to delete the model:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the model"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Model deleted successfully"})
}


