package main

import (
	"log"

	"github.com/teris-io/shortid"
)

var (
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
	ApiRouter().Run()
}
