package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/berabulut/capsule/api"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var sid *shortid.Shortid

type UserInput struct {
	URL string `bson:"url"`
}

func init() {
	var err error
	sid, err = shortid.New(1, shortid.DefaultABC, 232311234542)
	if err != nil {
		log.Fatal(err)
	}
	shortid.SetDefault(sid)

}

func main() {

	records, err := api.GetRecords()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		var request UserInput
		c.BindJSON(&request)
		shortid, _ := shortid.Generate()
		shortURL := &api.ShortURL{
			ID:              primitive.NewObjectID(),
			Key:             shortid,
			Value:           request.URL,
			CreatedAt:       time.Now(),
			LastTimeVisited: time.Now(),
			Clicks:          0,
		}
		api.ShortenURL(shortURL)
		records[shortid] = shortURL
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("URL to store: %v\n", shortURL),
		})
	})
	r.GET("/:key", func(c *gin.Context) {

		record, found := records[c.Request.URL.Path[1:]]

		if found {
			c.Redirect(http.StatusMovedPermanently, record.Value)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
