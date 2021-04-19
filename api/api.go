package api

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShortURL struct {
	ID              primitive.ObjectID `bson:"_id"`
	Key             string             `bson:"key"`
	Value           string             `bson:"value"`
	CreatedAt       time.Time          `bson:"created_at"`
	LastTimeVisited time.Time          `bson:"last_time_visited"`
	Clicks          int                `bson:"clicks"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("capsule").Collection("urls")
}

func ShortenURL(url *ShortURL) error {
	_, err := collection.InsertOne(ctx, url)
	return err
}
