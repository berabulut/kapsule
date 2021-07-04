package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func ApiRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://kapsule.click/"}

	r.Use(cors.New(config))

	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	r.POST("/shorten", ShortenURL())
	r.GET("/:key", GetDetails())
	r.GET("/details", GetMultipleRecords())

	return r
}
