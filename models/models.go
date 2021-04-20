package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortURL struct {
	ID              primitive.ObjectID `bson:"_id"`
	Key             string             `bson:"key"`
	Value           string             `bson:"value"`
	CreatedAt       time.Time          `bson:"created_at"`
	LastTimeVisited time.Time          `bson:"last_time_visited"`
	Clicks          int                `bson:"clicks"`
}

type UserInput struct {
	URL string `bson:"url"`
}
