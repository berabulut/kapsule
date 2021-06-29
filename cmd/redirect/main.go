package main

import (
	"log"
	"os"

	"github.com/teris-io/shortid"
)

var (
	sid   *shortid.Shortid
	Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
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

	RedirectRouter().Run(":8080")
}
