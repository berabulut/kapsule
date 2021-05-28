package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/berabulut/kapsule/helpers"
	db "github.com/berabulut/kapsule/mongo"
	"github.com/gin-gonic/gin"
)

var notFoundURL string

func init() {
	notFoundURL = os.Getenv("NOT_FOUND_URL") // godotenv should have loaded  by now
}

func RedirectURL() func(c *gin.Context) {

	return func(c *gin.Context) {
		key := c.Request.URL.Path[1:]

		record, err := db.GetRecord(key)

		if err != nil {
			c.Redirect(http.StatusSeeOther, notFoundURL)
			return
		}

		if record.Key != "" {
			userAgent := helpers.ParseUserAgent(c.Request.UserAgent())
			language := helpers.ParseLanguage(c.Request.Header.Get("Accept-Language"))
			countryCode, _ := helpers.GetCountryCode(c.Request.Header.Get("X-FORWARDED-FOR"))

			go c.HTML(http.StatusOK, "redirect.tmpl", gin.H{
				"title": record.Title,
				"url":   record.Value,
			})

			helpers.HandleClick(&record, userAgent, language, countryCode)

			err := db.HandleClick(key, record.Clicks, record.LastTimeVisited, record.Visits)
			if err != nil {
				log.Fatal(err)
			}

			return
		}

		c.Redirect(http.StatusSeeOther, notFoundURL)
	}
}
