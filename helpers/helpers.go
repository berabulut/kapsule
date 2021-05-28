package helpers

import (
	"strings"
	"time"

	"github.com/berabulut/kapsule/models"
	"github.com/mssola/user_agent"
)

func ParseUserAgent(userAgent string) models.UserAgent {
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

func ParseLanguage(language string) string {
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

func HandleClick(record *models.ShortURL, userAgent models.UserAgent, language string, countryCode string) { // update record's properties

	now := time.Now()
	index := len(record.Visits) - 1

	record.Clicks += 1

	if index >= 0 && sameDay(record.LastTimeVisited) { // new click within the same day

		record.LastTimeVisited = now

		visit := &record.Visits[index]
		visit.Clicks += 1
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
