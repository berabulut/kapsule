package api

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/berabulut/capsule/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func NewRecord(url *models.ShortURL) error {
	_, err := collection.InsertOne(ctx, url)
	return err
}

func GetRecords() (map[string]*models.ShortURL, error) {
	filter := bson.D{{}}
	return filterRecords(filter)
}

func filterRecords(filter interface{}) (map[string]*models.ShortURL, error) {
	records := make(map[string]*models.ShortURL)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return records, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var r models.ShortURL
		err := cur.Decode(&r)
		if err != nil {
			return records, err
		}

		records[r.Key] = &r
	}

	if err := cur.Err(); err != nil {
		return records, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(records) == 0 {
		return records, mongo.ErrNoDocuments
	}

	return records, nil

}
