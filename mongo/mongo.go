package api

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/berabulut/kapsule/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	err := godotenv.Load(".env") // tests can't find this one somehow
	if err != nil {
		err := godotenv.Load("../.env") // for tests
		if err != nil {
			log.Fatal(err)
		}
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

// func GetRecord(key string) (bson.M, error) {
// 	var record bson.M
// 	if err := collection.FindOne(ctx, bson.M{"key": key}).Decode(&record); err != nil {
// 		return nil, err
// 	}
// 	return record, nil
// }

func HandleClick(key string, clicks int, lastTimeVisited time.Time, visits []models.Visit) error {

	filter := bson.D{primitive.E{Key: "key", Value: key}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "clicks", Value: clicks},
		primitive.E{Key: "last_time_visited", Value: lastTimeVisited},
		primitive.E{Key: "visits", Value: visits},
	}}}

	r := &models.ShortURL{}

	return collection.FindOneAndUpdate(ctx, filter, update).Decode(r)
}

// 	t := &Task{}
// 	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
// }

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
