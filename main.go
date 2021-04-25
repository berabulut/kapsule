package main

import (
	"log"

	db "github.com/berabulut/capsule/mongo"
	"github.com/berabulut/capsule/routers"
	"github.com/teris-io/shortid"
	"golang.org/x/sync/errgroup"
)

var (
	g   errgroup.Group
	sid *shortid.Shortid
)

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

	g.Go(func() error {
		return routers.ApiRouter(records).Run(":8080")
	})

	g.Go(func() error {
		return routers.RedirectRouter(records).Run(":8081")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
