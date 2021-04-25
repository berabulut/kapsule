package routers

import (
	"github.com/berabulut/capsule/handlers"
	"github.com/berabulut/capsule/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRouter(records map[string]*models.ShortURL) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.POST("/shorten", handlers.ShortenURL(records))

	return r
}

func RedirectRouter(records map[string]*models.ShortURL) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/:key", handlers.RedirectURL(records))

	return r

}
