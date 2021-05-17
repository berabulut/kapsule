package routers

import (
	"net/http"
	"strings"
	"time"

	"github.com/berabulut/kapsule/helpers"
	"github.com/berabulut/kapsule/models"
	db "github.com/berabulut/kapsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShortenURL(records map[string]*models.ShortURL) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request models.UserInput
		c.BindJSON(&request)
		shortid, _ := shortid.Generate()

		resp, err := http.Get(request.URL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		htmlTitle, ok := helpers.GetHtmlTitle(resp.Body)

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
		}

		db.NewRecord(shortURL)
		records[shortid] = shortURL

		c.JSON(200, gin.H{
			"id":    shortURL.Key,
			"title": shortURL.Title,
		})
	}
}

func GetDetails(records map[string]*models.ShortURL) func(c *gin.Context) {
	return func(c *gin.Context) {

		key := c.Request.URL.Path[1:]
		record, found := records[key]

		if found {
			c.JSON(200, gin.H{
				"record": record,
			})
			return
		}

		c.JSON(404, gin.H{
			"record": "",
		})

	}
}

func GetMultipleRecords(records map[string]*models.ShortURL) func(c *gin.Context) {
	return func(c *gin.Context) {
		keys := c.QueryMap("keys")
		var values []models.ShortURL

		for _, key := range keys {
			record, found := records[key]
			if found {
				values = append(values, *record)
			}
		}

		if len(values) > 0 {
			c.JSON(200, gin.H{
				"records": values,
			})
			return
		}

		c.JSON(404, gin.H{
			"records": nil,
		})

	}
}
