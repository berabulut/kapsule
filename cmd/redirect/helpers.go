package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/berabulut/kapsule/models"
	"github.com/mssola/user_agent"
)

type response struct {
	CountryCode string `json:"country_code3"`
}

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

func GetCountryCode(IP string) (string, error) {

	url := fmt.Sprintf("https://get.geojs.io/v1/ip/geo/%s.json", IP)
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return "", err
	}

	resp := response{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return "", err
	}

	return resp.CountryCode, nil
}

func SameDay(lastTime time.Time) bool {
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
