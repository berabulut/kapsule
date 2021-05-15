package handlers

import (
	"github.com/berabulut/kapsule/models"
	"github.com/mssola/user_agent"
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
