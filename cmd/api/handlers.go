package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/berabulut/kapsule/models"
	db "github.com/berabulut/kapsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInput struct {
	URL            string `json:"url"`
	OptionsEnabled bool   `json:"options_enabled"`
	Duration       int    `json:"duration"`
	Message        string `json:"message"`
}

func ShortenURL() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request UserInput
		c.BindJSON(&request)

		shortid, err := shortid.Generate()
		if err != nil {
			Error.Println(err)
			c.JSON(500, gin.H{})
			return
		}

		// to get title of html page
		resp, err := http.Get(request.URL)
		if err != nil {
			Error.Println(err)
		}
		defer resp.Body.Close()

		htmlTitle, ok := GetHtmlTitle(resp.Body)
		if !ok {
			htmlTitle = "Not found"
		}
		htmlTitle = strings.TrimSpace(htmlTitle)

		shortURL := &models.ShortURL{
			ID:        primitive.NewObjectID(),
			Key:       shortid,
			Value:     request.URL,
			CreatedAt: time.Now().Unix(),
			Clicks:    0,
			Visits:    []models.Visit{},
			Title:     htmlTitle,
			Options: models.Options{
				Enabled:  request.OptionsEnabled,
				Duration: request.Duration % 11,
				Message:  request.Message,
			},
		}

		db.NewRecord(shortURL)

		c.JSON(200, gin.H{
			"id":    shortURL.Key,
			"title": shortURL.Title,
		})
	}
}

func GetDetails() func(c *gin.Context) {
	return func(c *gin.Context) {
		key := c.Request.URL.Path[1:]

		if record, _ := db.GetRecord(key); record.Key != "" {
			c.JSON(200, gin.H{
				"record": record,
			})
			return
		}

		c.JSON(404, gin.H{})

	}
}

func GetMultipleRecords() func(c *gin.Context) {
	return func(c *gin.Context) {

		keys := c.QueryMap("keys")

		var values []models.ShortURL

		for _, key := range keys {

			if record, _ := db.GetRecord(key); record.Key != "" {
				values = append(values, record)
			}

		}

		if len(values) > 0 {
			c.JSON(200, gin.H{
				"records": values,
			})
			return
		}

		c.JSON(404, gin.H{})

	}
}
