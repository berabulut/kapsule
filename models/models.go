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

type Header struct {
	AcceptLanguage string `json:"Accept-Language"`
	UserAgent      string `json:"User-Agent"`
}

type Visit struct {
	Clicks    int
	Date      int64       // unix timestamp
	UserAgent []UserAgent `bson:"user_agent"`
	Language  []string    `bson:"language"`
}

type ShortURL struct {
	ID              primitive.ObjectID `bson:"_id"`
	Key             string             `bson:"key"`
	Value           string             `bson:"value"`
	CreatedAt       int64              `bson:"created_at"` // unix timestamp
	LastTimeVisited time.Time          `bson:"last_time_visited"`
	Clicks          int                `bson:"clicks"`
	Visits          []Visit            `bson:"visits"`
}

type UserInput struct {
	URL string `bson:"url"`
}
