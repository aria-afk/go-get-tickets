package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

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
	p, _ := pg.NewPG()
	r.GET("/vendors", func(c *gin.Context) {
		// type Vendor struct {
		// 	UUID      string
		// 	Name      string
		// 	OwnerUUID string
		// 	CreatedAt time.Time
		// 	UpdatedAt time.Time
		// }
		// vendor := Vendor{}
		var UUID string
		// p.Conn.QueryRow("SELECT uuid,name,owner_uuid,created_at,updated_at FROM vendors LIMIT 1").Scan(&UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt)
		err := p.Conn.QueryRow("SELECT uuid FROM vendors LIMIT 1").Scan(&UUID)
		// fmt.Printf("%#v\n", vendor)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print(UUID)

		// TODO: check if i got a no row error

		// i would have a field in gin.H that i put the struct in maybe is what aria says and i don't marshall it? if i want to nest it
		// j, _ := json.Marshal(vendor)
		// j, _ := json.Marshal(UUID)

		// TODO: this is not returning what i want, try using gin.H
		c.JSON(http.StatusOK, UUID)
	})
	r.GET("/vendors/:vendor_uuid", func(c *gin.Context) {
		type Vendor struct {
			UUID      string
			Name      string
			OwnerUUID string
			CreatedAt time.Time
			UpdatedAt time.Time
		}
		vendor := Vendor{}
		vu := c.Param("vendor_uuid")
		// p.Conn.QueryRow("SELECT uuid,name,owner_uuid,created_at,updated_at FROM vendors LIMIT 1").Scan(&UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt)
		err := p.Conn.QueryRow("SELECT uuid, name, owner_uuid, created_at, updated_at FROM vendors WHERE uuid = '"+vu+"' LIMIT 1").
			Scan(&vendor.UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", vendor)
		fmt.Print(vendor)

		// TODO: check if i got a no row error
		// TODO: load some example data in for tests

		// i would have a field in gin.H that i put the struct in maybe is what aria says and i don't marshall it? if i want to nest it
		// j, _ := json.Marshal(vendor)
		// j, _ := json.Marshal(vendor)

		c.JSON(http.StatusOK, vendor)
	})
	//
	// GET vendor
	// PUT vendor
	// PATCH vendor
	// DELETE vendor
	// cookies? or JWT?
	// GET user for viewing your profile or settings or name??
	// PUT user
	// PATCH user
	// DELETE user
	// GET event
	// PUT event
	// provision tickets at event creation directly into db (means there needs to be upper limit)
	// that is also where the qr codes get generated
	// PATCH event to make edits or to hide or show an event
	// DELETE event
	// for image upload we will try to use presigned urls
	// GET presigned url (image name as param) for uploading to object store
	// GET ticket
	// PUT ticket
	// PATCH ticket
	// DELETE ticket (for refund)
	// GET receipt
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
