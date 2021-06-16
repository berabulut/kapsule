package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAgent struct {
	Mobile         bool
	Platform       string
	OS             string
	Browser        string
	BrowserVersion string
}

type Visit struct {
	Clicks      int
	Date        int64 // unix timestamp
	UserAgent   []UserAgent
	Language    []string
	CountryCode []string
}

type Options struct {
	Enabled  bool
	Duration int
	Note     string
}

type ShortURL struct {
	ID              primitive.ObjectID `bson:"_id"`
	Key             string             `bson:"key"`
	Value           string             `bson:"value"`
	Title           string             `bson:"title"`
	CreatedAt       int64              `bson:"created_at"` // unix timestamp
	LastTimeVisited time.Time          `bson:"last_time_visited"`
	Clicks          int                `bson:"clicks"`
	Visits          []Visit            `bson:"visits"`
	Options         Options            `bson:"options"`
}

type UserInput struct {
	URL            string `json:"url"`
	OptionsEnabled bool   `json:"options_enabled"`
	Duration       int    `json:"duration"`
	Note           string `json:"note"`
}
