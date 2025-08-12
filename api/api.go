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

func GetVendorUsers(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []VendorUser
		rows, err := p.Conn.Query("SELECT uuid, name, vendor_uuid, permissions, email, created_at, updated_at FROM vendor_users")
		// TODO handle no rows error or something
		defer rows.Close()
		for rows.Next() {
			var user VendorUser
			if err := rows.Scan(&user.UUID, &user.Name, &user.VendorUUID, &user.Permissions, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Printf("Error scanning vendor_users: %v\n", err)
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

func PutVendorUsers(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user VendorUser
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		_, err := p.Conn.Query(
			"INSERT INTO vendor_users(name, vendor_uuid, permissions, email) VALUES($1, $2, $3, $4)",
			user.Name, user.VendorUUID, user.Permissions, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusAccepted, gin.H{})
	}
}

func GetVendors(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vendors []Vendor
		rows, err := p.Conn.Query("SELECT uuid, name, created_at, updated_at FROM vendors")
		defer rows.Close()
		for rows.Next() {
			var vendor Vendor
			if err := rows.Scan(&vendor.UUID, &vendor.Name, &vendor.CreatedAt, &vendor.UpdatedAt); err != nil {
				log.Printf("Error scanning vendors: %v\n", err)
			} else {
				vendors = append(vendors, vendor)
			}
		}
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

func SetupRoutes(p *pg.PG) *gin.Engine {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	r := gin.Default()

	r.GET("/vendor_users", GetVendorUsers(p))
	r.PUT("/vendor_users", PutVendorUsers(p))

	r.GET("/vendors", GetVendors(p))
	r.GET("/vendors/:vendor_uuid", func(c *gin.Context) {
		vendor := Vendor{}
		vu := c.Param("vendor_uuid")
		fmt.Println(vu)
		// p.Conn.QueryRow("SELECT uuid,name,owner_uuid,created_at,updated_at FROM vendors LIMIT 1").Scan(&UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt)
		// err := p.Conn.QueryRow("SELECT uuid, name, owner_uuid, created_at, updated_at FROM vendors WHERE uuid = ? LIMIT 1", vu).
		//		Scan(&vendor.UUID, &vendor.Name, &vendor.OwnerUUID, &vendor.CreatedAt, &vendor.UpdatedAt)
		// err := p.Conn.QueryRow("SELECT uuid, name, created_at, updated_at FROM vendors WHERE uuid = ? LIMIT 1", vu).
		err := p.Conn.QueryRow("SELECT uuid, name, created_at, updated_at FROM vendors WHERE uuid = $1 LIMIT 1", vu).
			Scan(&vendor.UUID, &vendor.Name, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%#v\n", vendor)
		fmt.Print(vendor)

		// TODO: check if i got a no row error
	})

	return r
}

func ServeAPI() {
	p, _ := pg.NewPG()
	r := SetupRoutes(p)
	r.Run()
}
