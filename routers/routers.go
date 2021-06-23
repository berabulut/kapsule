package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/chenjiandongx/ginprom"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ApiRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	// prometheus middleware
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

func RedirectRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	// prometheus middleware
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	//r.LoadHTMLGlob("templates/**")
	r.LoadHTMLGlob("./web/templates/**")
	r.Static("/static", "./web/static")

	r.GET("/:key", RedirectURL())

	return r

}
