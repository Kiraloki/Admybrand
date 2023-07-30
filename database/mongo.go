package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace the connection string with your MongoDB URL
	clientOptions := options.Client().ApplyURI("mongodb+srv://lightyagami81161:J9CdXbZSoJMDOgTD@cluster0.q7grb1t.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the MongoDB to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}  

	



	log.Println("Connected to MongoDB successfully!")
}

func GetMongoDBClient() *mongo.Client {
	return client
}
