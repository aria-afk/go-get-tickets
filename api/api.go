package api

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/aria-afk/go-get-tickets/pg"
)

type Vendor struct {
	UUID      string    `json:"uuid,omitempty"`
	Name      string    `json:"name,omitempty"`
	OwnerUUID string    `json:"owner_uuid,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type VendorUser struct {
	UUID        string    `json:"uuid,omitempty"`
	Name        string    `json:"name,omitempty"`
	VendorUUID  string    `json:"vendor_uuid,omitempty"`
	Permissions string    `json:"permissions,omitempty"`
	Email       string    `json:"email,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
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

func PutVendorUser(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user VendorUser
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err := p.Conn.Query(
			"INSERT INTO vendor_users(name, vendor_uuid, permissions, email) VALUES($1, $2, $3, $4)",
			user.Name, user.VendorUUID, user.Permissions, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			log.Println(err)
		} else {
			c.JSON(http.StatusAccepted, gin.H{})
		}
	}
}

func DeleteVendorUser(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := p.Conn.Query(
			"DELETE FROM vendor_users WHERE uuid = $1", c.Param("user_uuid"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

func GetVendorUser(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user VendorUser
		vu := c.Param("user_uuid")
		err := p.Conn.QueryRow("SELECT uuid, name, vendor_uuid, permissions, email, created_at, updated_at FROM vendors WHERE uuid = $1 LIMIT 1", vu).
			Scan(&user.UUID, &user.Name, &user.VendorUUID, &user.Permissions, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, user)
		}
		// fmt.Printf("%#v\n", vendor)
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

func PutVendor(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vendor Vendor
		if err := c.BindJSON(&vendor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err := p.Conn.Query(
			"INSERT INTO vendor_users(name, owner_uuid) VALUES($1, $2)",
			vendor.Name, vendor.OwnerUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusAccepted, gin.H{})
		}
	}
}

func DeleteVendor(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := p.Conn.Query(
			"DELETE FROM vendors WHERE uuid = $1", c.Param("vendor_uuid"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

func GetVendor(p *pg.PG) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vendor Vendor
		vu := c.Param("vendor_uuid")
		err := p.Conn.QueryRow("SELECT uuid, name, created_at, updated_at FROM vendors WHERE uuid = $1 LIMIT 1", vu).
			Scan(&vendor.UUID, &vendor.Name, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, vendor)
		}
		// fmt.Printf("%#v\n", vendor)
	}
}

func SetupRoutes(p *pg.PG) *gin.Engine {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	r := gin.Default()

	r.GET("/vendor_users", GetVendorUsers(p))
	r.PUT("/vendor_users", PutVendorUser(p))
	r.GET("/vendor_users/:user_uuid", GetVendorUser(p))
	r.DELETE("/vendor_users/:user_uuid", DeleteVendorUser(p))

	r.GET("/vendors", GetVendors(p))
	r.PUT("/vendors", PutVendor(p))
	r.GET("/vendors/:vendor_uuid", GetVendor(p))
	r.DELETE("/vendors/:vendor_uuid", DeleteVendor(p))

	return r
}

func ServeAPI() {
	p, _ := pg.NewPG()
	r := SetupRoutes(p)
	r.Run()
}
