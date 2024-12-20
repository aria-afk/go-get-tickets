package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/aria-afk/go-get-tickets/pg"
	"github.com/aria-afk/go-get-tickets/utils"
)

func main() {
	utils.LoadEnv("dev.env")

	// flags
	migrate := flag.String("migrate", "", "Type of migration to run, valid options are [up/down]")
	flag.Parse()

	if *migrate != "" {
		fmt.Printf("Running migration %s\n", *migrate)
		db, err := pg.NewPG()
		if err != nil {
			fmt.Printf("Error getting DB connection:\n%s", err)
			os.Exit(1)
		}
		err = db.Migrate(*migrate)
		if err != nil {
			os.Exit(1)
		}
	}

	r := gin.Default()
	// PUT vendor
	// GET vendor
	// PATCH vendor
	// DELETE vendor
	// PUT user
	// cookies? or JWT?
	// GET user for viewing your profile or settings or name??
	// PATCH user
	// DELETE user
	// PUT event
	// provision tickets at event creation directly into db (means there needs to be upper limit)
	// that is also where the qr codes get generated
	// PATCH event to make edits or to hide or show an event
	// GET event
	// DELETE event
	// for image upload we will try to use presigned urls
	// GET presigned url (image name as param) for uploading to object store
	// dont need a PUT for ticket yet
	// GET ticket
	// PATCH ticket
	// DELETE ticket (for refund)
	// PUT receipt
	// POST refund
	// POST purchase
	// vendors will need to provide paypal or stripe account info
	// we'll say they need to give us paypal api key i think
	// TODO research this payments thing more. issue is marketplace solutions arent free.
	// TODO we should put docs including screenshots of a phone into the ui
	//
	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "outer ontext",
		})
	})
	r.Run()
}
