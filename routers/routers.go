package routers

import (
	"github.com/berabulut/kapsule/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRouter(records map[string]*models.ShortURL) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	r.POST("/shorten", ShortenURL(records))
	r.GET("/:key", GetDetails(records))
	r.GET("/details", GetMultipleRecords(records))
	// r.GET("/:key", db.GetRecord("n1K1N6bK2"))

	return r
}

func RedirectRouter(records map[string]*models.ShortURL) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(gin.Logger())

	r.GET("/:key", RedirectURL(records))

	return r

}
