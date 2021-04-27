package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/berabulut/capsule/models"
	db "github.com/berabulut/capsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ShortenURL(records map[string]*models.ShortURL) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request models.UserInput
		c.BindJSON(&request)
		shortid, _ := shortid.Generate()
		shortURL := &models.ShortURL{
			ID:        primitive.NewObjectID(),
			Key:       shortid,
			Value:     request.URL,
			CreatedAt: time.Now().Unix(),
			Clicks:    0,
			Visits:    []models.Visit{},
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
			userAgent := parseUserAgent(c.Request.UserAgent())
			language := parseLanguage(c.Request.Header.Get("Accept-Language"))
			HandleClick(records[key], userAgent, language)

			c.JSON(200, gin.H{
				"record": records[key],
				// "request":   c.Request.UserAgent(),
				// "request1":  c.Request.Header,
				// "request2":  c.Request.Header.Get("Accept-Language"),
				// "request12": c.Request.RemoteAddr,
			})
			err := db.HandleClick(*records[key])
			if err != nil {
				log.Fatal(err)
			}
			c.Redirect(http.StatusFound, records[key].Value)
		} else {
			c.Redirect(http.StatusSeeOther, "http://localhost:3000/")
		}
	}
}

func HandleClick(record *models.ShortURL, userAgent models.UserAgent, language string) {

	now := time.Now()
	index := len(record.Visits) - 1

	record.Clicks += 1

	if index >= 0 && sameDay(record.LastTimeVisited) { // new click withing the same day
		record.LastTimeVisited = now

		visit := record.Visits[index]
		visit.Clicks += 1
		visit.Language = append(visit.Language, language)
		visit.UserAgent = append(visit.UserAgent, userAgent)

		record.Visits[index] = visit // if it's not a new day it must be the last element of slice

		return
	}

	record.LastTimeVisited = now
	record.Visits = append(record.Visits, models.Visit{
		Clicks:    1,
		Date:      time.Now().Unix(),
		Language:  []string{language},
		UserAgent: []models.UserAgent{userAgent},
	})
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

func parseUserAgent(userAgent string) models.UserAgent {
	ua := user_agent.New(userAgent)

	browser, browserVersion := ua.Browser()

	return models.UserAgent{
		Mobile:         ua.Mobile(),
		Platform:       ua.Platform(),
		OS:             ua.OS(),
		Browser:        browser,
		BrowserVersion: browserVersion,
	}
}

func parseLanguage(language string) string {
	if len(language) > 0 {
		i := strings.Index(language, ",")
		return language[:i]
	}
	return ""
}
