package main

import (
	"log"

	"github.com/berabulut/kapsule/routers"
	"github.com/teris-io/shortid"
	"golang.org/x/sync/errgroup"
)

var (
	g   errgroup.Group
	sid *shortid.Shortid
)

func init() {
	var err error
	const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_*"
	sid, err = shortid.New(1, alphabet, 232311234542)
	if err != nil {
		log.Fatal(err)
	}
	shortid.SetDefault(sid)

}

func main() {

	g.Go(func() error {
		return routers.ApiRouter().Run(":8080")
	})

	g.Go(func() error {
		return routers.RedirectRouter().Run(":8081")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
