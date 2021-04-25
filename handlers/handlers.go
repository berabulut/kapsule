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
		shortURL := &models.ShortURL{
			ID:              primitive.NewObjectID(),
			Key:             shortid,
			Value:           request.URL,
			CreatedAt:       time.Now(),
			LastTimeVisited: time.Now(),
			Clicks:          0,
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
		record, found := records[c.Request.URL.Path[1:]]

		if found {
			c.Redirect(http.StatusMovedPermanently, record.Value)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}
