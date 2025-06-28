package api

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aria-afk/go-get-tickets/pg"
)

type Vendor struct {
	UUID      string
	Name      string
	OwnerUUID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VendorUser struct {
	UUID        string
	Name        string
	VendorUUID  string
	Permissions string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetVendorUsersHandler(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []VendorUser
		rows, err := p.Conn.Query("SELECT uuid, name, vendor_uuid, permissions, email, created_at, updated_at FROM vendor_users")
		defer rows.Close()
		for rows.Next() {
			var user VendorUser
			if err := rows.Scan(&user.UUID, &user.Name, &user.VendorUUID, &user.Permissions, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Println("Error scanning vendor_users: %v", err)
			} else {
				users = append(users, user)
			}
		}
		if err = rows.Err(); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}

func GetVendorsHandler(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vendors []Vendor
		// rows, err := p.Conn.Query("SELECT uuid, name, owner_uuid, created_at, updated_at FROM vendors")
		rows, err := p.Conn.Query("SELECT uuid, name, created_at, updated_at FROM vendors")
		defer rows.Close()
		for rows.Next() {
			var vendor Vendor
			// if err := rows.Scan(&vendor.UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt); err != nil {
			if err := rows.Scan(&vendor.UUID, &vendor.Name, &vendor.CreatedAt, &vendor.UpdatedAt); err != nil {
				log.Println("Error scanning vendors: %v", err)
			} else {
				vendors = append(vendors, vendor)
			}
		}
		// fmt.Printf("%#v\n", vendor)
		if err = rows.Err(); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, vendors)
		}
	}
}

func GetVendorByUUIDHandler(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		vendor := Vendor{}
		vu := c.Param("vendor_uuid")
		fmt.Println(vu)
		err := p.Conn.QueryRow("SELECT uuid, name, created_at, updated_at FROM vendors WHERE uuid = $1 LIMIT 1", vu).
			Scan(&vendor.UUID, &vendor.Name, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", vendor)
		fmt.Print(vendor)

		// TODO: check if i got a no row error
		// TODO: load some example data in for tests

		c.JSON(http.StatusOK, vendor)
	}
}

func ServeAPI() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	r := gin.Default()
	p, _ := pg.NewPG()

	r.GET("/vendor_users", GetVendorUsersHandler(p))
	r.PUT("/vendor_users", func(c *gin.Context) {
	})

	r.GET("/vendors", GetVendorsHandler(p))
	r.GET("/vendors/:vendor_uuid", GetVendorByUUIDHandler(p))

	// GET vendor
	// PUT vendor
	// TODO use WITH INSERT INTO ... RETURNING uuid AS ... to insert in a transaction?
	// PATCH vendor
	// DELETE vendor
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
	// DELETE ticket (for refund?)
	// GET receipt
	// PUT receipt
	// POST refund
	// POST purchase
	//
	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "outer ontext",
		})
	})
	r.Run()
}
