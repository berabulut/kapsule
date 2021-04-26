package handlers

import (
	"net/http"
	"time"

	"github.com/berabulut/capsule/models"
	db "github.com/berabulut/capsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShortenURL(records map[string]*models.ShortURL) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request models.UserInput
		c.BindJSON(&request)
		shortid, _ := shortid.Generate()
		now := time.Now().Unix()
		shortURL := &models.ShortURL{
			ID:              primitive.NewObjectID(),
			Key:             shortid,
			Value:           request.URL,
			CreatedAt:       now,
			LastTimeVisited: time.Now(),
			Clicks:          0,
			Visits:          []models.Visit{},
		}
		db.NewRecord(shortURL)
		records[shortid] = shortURL
		c.JSON(200, gin.H{
			"id": shortURL.Key,
		})
	}
}

func RedirectURL(records map[string]*models.ShortURL) func(c *gin.Context) {

	return func(c *gin.Context) {
		key := c.Request.URL.Path[1:]
		_, found := records[key]

		if found {
			record := HandleClick(records[key])
			db.HandleClick(record)
			records[key] = record
			c.Redirect(http.StatusMovedPermanently, record.Value)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

func HandleClick(record *models.ShortURL) *models.ShortURL {
	length := len(record.Visits)

	if length > 0 && sameDay(record.LastTimeVisited) { // new click withing the same day
		record.Visits[len(record.Visits)-1].Clicks += 1 // if it's not a new day it must be the last element of slice
		return record
	}

	record.Visits = append(record.Visits, models.Visit{
		Clicks: 1,
		Date:   time.Now().Unix(),
	})

	record.Clicks += 1
	record.LastTimeVisited = time.Now()

	return record
}

func sameDay(lastTime time.Time) bool {
	now := time.Now()
	// now := time.Date(
	// 	2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	if lastTime.Day() != now.Day() {
		return false
	}

	if lastTime.Month() != now.Month() {
		return false
	}

	if lastTime.Year() != now.Year() {
		return false
	}

	return true
}
