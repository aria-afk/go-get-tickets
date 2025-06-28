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

func NotImplementedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, nil)
	}
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
	r.GET("/vendors", GetVendorsHandler(p))
	r.GET("/vendors/:vendor_uuid", GetVendorByUUIDHandler(p))

	// TODO: Implement
	r.GET("/vendor_users/:user_uuid", NotImplementedHandler())
	r.PATCH("/vendor_users/:user_uuid", NotImplementedHandler())
	r.DELETE("/vendor_users/:user_uuid", NotImplementedHandler())
	r.PUT("/vendor_users", NotImplementedHandler())

	r.PATCH("/vendors/:vendor_uuid", NotImplementedHandler())
	r.DELETE("/vendors/:vendor_uuid", NotImplementedHandler())
	r.PUT("/vendors", NotImplementedHandler())

	r.GET("/events", NotImplementedHandler())
	r.GET("/events/:event_uuid", NotImplementedHandler())
	r.PUT("/events", NotImplementedHandler())
	r.PATCH("/events/:event_uuid", NotImplementedHandler())
	r.DELETE("/events/:event_uuid", NotImplementedHandler())

	r.GET("/tickets", NotImplementedHandler())
	r.GET("/tickets/:ticket_uuid", NotImplementedHandler())
	r.PUT("/tickets", NotImplementedHandler())
	r.PATCH("/tickets/:ticket_uuid", NotImplementedHandler())
	r.DELETE("/tickets/:ticket_uuid", NotImplementedHandler())

	// for image upload we will try to use presigned urls
	// GET presigned url (image name as param) for uploading to object store
	// GET receipt
	// PUT receipt
	// POST refund
	// POST purchase

	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "outer ontext",
		})
	})
	r.Run()
}
