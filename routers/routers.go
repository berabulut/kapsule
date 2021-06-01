package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

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

	//r.LoadHTMLGlob("templates/**")
	r.LoadHTMLGlob("./web/templates/**")
	r.Static("/static", "./web/static")

	r.GET("/:key", RedirectURL())

	return r

}
