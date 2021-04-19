package main

import (
	"fmt"
	"log"
	"time"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var sid *shortid.Shortid

type URL struct {
	ID              primitive.ObjectID `bson:"_id"`
	CreatedAt       time.Time          `bson:"created_at"`
	LastTimeVisited time.Time          `bson:"last_time_visited"`
	Value           string             `bson:"value"`
	Clicks          int                `bson:"clicks"`
}

func init() {
	var err error
	sid, err = shortid.New(1, shortid.DefaultABC, 232311234542)
	if err != nil {
		log.Fatal(err)
	}
	shortid.SetDefault(sid)
}

func main() {
	fmt.Printf(shortid.Generate())

}
