package main

import (
	"net/http"
	"os"
	"time"

	"github.com/berabulut/kapsule/models"
	db "github.com/berabulut/kapsule/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

var notFoundURL = os.Getenv("NOT_FOUND_URL")

// update record properties
func handleClick(record *models.ShortURL, userAgent models.UserAgent, language string, countryCode string) {

	now := time.Now()
	index := len(record.Visits) - 1

	record.Clicks++

	// new click within the same day
	if index >= 0 && SameDay(record.LastTimeVisited) {

		record.LastTimeVisited = now

		visit := &record.Visits[index]
		visit.Clicks++
		visit.Language = append(visit.Language, language)
		visit.UserAgent = append(visit.UserAgent, userAgent)
		visit.CountryCode = append(visit.CountryCode, countryCode)

		return
	}

	record.LastTimeVisited = now
	record.Visits = append(record.Visits, models.Visit{
		Clicks:      1,
		Date:        time.Now().Unix(),
		Language:    []string{language},
		UserAgent:   []models.UserAgent{userAgent},
		CountryCode: []string{countryCode},
	})
}

func RedirectURL() func(c *gin.Context) {

	return func(c *gin.Context) {

		key := c.Request.URL.Path[1:]

		// if record has been found
		if record, err := db.GetRecord(key); record.Key != "" {
			userAgent := ParseUserAgent(c.Request.UserAgent())
			language := ParseLanguage(c.Request.Header.Get("Accept-Language"))
			countryCode, _ := GetCountryCode(c.Request.Header.Get("X-FORWARDED-FOR"))

			// show a static page before redirecting
			if record.Options.Enabled {
				c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
					"title":    record.Title,
					"url":      record.Value,
					"duration": record.Options.Duration,
					"message":  record.Options.Message,
				})
			} else {
				c.Redirect(http.StatusFound, record.Value)
			}

			// update record properties before updating database
			handleClick(&record, userAgent, language, countryCode)

			// update db with returned values
			if err := db.HandleClick(key, record.Clicks, record.LastTimeVisited, record.Visits); err != nil {
				Error.Println(err)
			}

			return

		} else if err != nil {
			Error.Println(err)
		}

		c.Redirect(http.StatusSeeOther, notFoundURL)
	}
}

func RedirectRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	//r.LoadHTMLGlob("templates/**")
	r.LoadHTMLGlob("../../web/templates/**")
	r.Static("/static", "../../web/static")

	r.GET("/:key", RedirectURL())

	return r

}
