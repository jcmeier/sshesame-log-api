package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	database mongo.Database
	client   mongo.Client
}

func CreateRepository(mongodbURI string) MongoRepository {
	clientOptions := options.Client().ApplyURI(mongodbURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database("logs")
	return MongoRepository{
		database: *database,
		client:   *client,
	}

}

func (repo *MongoRepository) FindAll() []map[string]interface{} {
	collection := repo.database.Collection("attacks")

	sortOptions := options.Find().SetSort(bson.D{{"lastUpdate", -1}}).SetLimit(100)

	// Execute the query with sort options.
	cursor, err := collection.Find(context.Background(), bson.D{}, sortOptions)
	if err != nil {
		log.Println("Find did not return someting: ", err)
		return nil
	}

	var results []map[string]interface{}
	if err := cursor.All(context.Background(), &results); err != nil {
		log.Println("Error in cursor: ", err)
		return nil
	}

	return results
}

func (repo *MongoRepository) Disconnect() {
	if err := repo.client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
