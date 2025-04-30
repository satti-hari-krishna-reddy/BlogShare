package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection
var cacheCollection *mongo.Collection
var scheduledItemsCollection *mongo.Collection

func InitMongoDb() {
	mongoURI := os.Getenv("MONGO_URI")
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB")

	if mongoHost == "" {
		mongoHost = "localhost:27017"
	}

	if mongoURI == "" {
		if mongoUser != "" && mongoPassword != "" {
			mongoURI = fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPassword, mongoHost)
		} else {
			log.Println("[WARN] No MONGO_URI and no auth credentials, connecting without auth")
			mongoURI = fmt.Sprintf("mongodb://%s", mongoHost)
		}
	}

	if dbName == "" {
		dbName = "social-scribe"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("[ERROR] Failed connecting to MongoDB:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("[ERROR] Could not ping MongoDB:", err)
	}

	log.Printf("[INFO] Successfully connected to MongoDB (%s), using database %q\n", mongoURI, dbName)

	userCollection = client.Database(dbName).Collection("users")
	cacheCollection = client.Database(dbName).Collection("cache")
	scheduledItemsCollection = client.Database(dbName).Collection("scheduled_items")

	err = CreateIndexes()
	if err != nil {
		log.Println("[ERROR] Failed creating indexes:", err)
	}
	log.Println("[INFO] Successfully connected to MongoDB")
}

func CreateIndexes() error {
	ctx := context.TODO()

	indexes := []mongo.IndexModel{
		// TTL index for automatic expiration
		{
			Keys:    bson.D{{Key: "expiresAt", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
		// Unique index for cache keys
		{
			Keys:    bson.D{{Key: "key", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := cacheCollection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("[ERROR] Error creating indexes: %v", err)
		return err
	}

	log.Println("[INFO] Successfully created indexes for cache collection")
	return nil
}
