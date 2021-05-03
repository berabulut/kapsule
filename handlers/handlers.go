package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/berabulut/capsule/handlers/title"
	"github.com/berabulut/capsule/models"
	db "github.com/berabulut/capsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func handleClick(record *models.ShortURL, userAgent models.UserAgent, language string, remoteAddr string, xForwardedFor string) {

	now := time.Now()
	index := len(record.Visits) - 1

	record.Clicks += 1

	if index >= 0 && sameDay(record.LastTimeVisited) { // new click withing the same day

		record.LastTimeVisited = now

		visit := &record.Visits[index]
		visit.Clicks += 1
		visit.Language = append(visit.Language, language)
		visit.UserAgent = append(visit.UserAgent, userAgent)
		visit.RemoteAddr = remoteAddr
		visit.XForwardedFor = xForwardedFor

		return
	}

	record.LastTimeVisited = now
	record.Visits = append(record.Visits, models.Visit{
		Clicks:        1,
		Date:          time.Now().Unix(),
		Language:      []string{language},
		UserAgent:     []models.UserAgent{userAgent},
		RemoteAddr:    remoteAddr,
		XForwardedFor: xForwardedFor,
	})
}

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

		htmlTitle, ok := title.GetHtmlTitle(resp.Body)

		if !ok {
			htmlTitle = "Not found"
		}

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

func RedirectURL(records map[string]*models.ShortURL) func(c *gin.Context) {

	return func(c *gin.Context) {
		key := c.Request.URL.Path[1:]

		record, found := records[key]

		if found {
			userAgent := parseUserAgent(c.Request.UserAgent())
			language := parseLanguage(c.Request.Header.Get("Accept-Language"))
			remoteAddr := c.Request.RemoteAddr
			xForwardedFor := c.Request.Header.Get("X-FORWARDED-FOR")

			handleClick(record, userAgent, language, remoteAddr, xForwardedFor)

			err := db.HandleClick(key, record.Clicks, record.LastTimeVisited, record.Visits)
			if err != nil {
				log.Fatal(err)
			}
			c.Redirect(http.StatusFound, records[key].Value)
		} else {
			c.Redirect(http.StatusSeeOther, "http://localhost:3000/")
		}
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
