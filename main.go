package main

import (
	"log"

	"github.com/berabulut/capsule/handlers"
	db "github.com/berabulut/capsule/mongo"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

var sid *shortid.Shortid

func init() {
	var err error
	sid, err = shortid.New(1, shortid.DefaultABC, 232311234542)
	if err != nil {
		log.Fatal(err)
	}
	shortid.SetDefault(sid)

}

func main() {

	records, err := db.GetRecords()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/shorten", handlers.ShortenURL(records))
	r.GET("/:key", handlers.Redirect(records))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
