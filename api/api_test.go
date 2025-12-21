package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aria-afk/go-get-tickets/pg"
	"github.com/stretchr/testify/assert"
)

func TestGetVendorUsers(t *testing.T) {
	p, _ := pg.NewPG()
	r := SetupRoutes(p)

	req, _ := http.NewRequest("GET", "/vendor_users", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPutVendorUserShouldSucceed(t *testing.T) {
	p, _ := pg.NewPG()
	r := SetupRoutes(p)

	u := VendorUser{
		Name:        "test User1",
		Permissions: "admin",
		Email:       "test@localhost",
		// VendorUUID:  "aaaaa0",
	}
	u_json, _ := json.Marshal(u)
	fmt.Println(string(u_json))
	req, _ := http.NewRequest("PUT", "/vendor_users", bytes.NewBuffer(u_json))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestPatchVendorUserEmailShouldSucceed(t *testing.T) {
	p, _ := pg.NewPG()
	r := SetupRoutes(p)

	u_init := VendorUser{
		Name:        "test User2",
		Permissions: "admin",
		Email:       "test2@localhost",
	}
	u_init_json, _ := json.Marshal(u_init)
	req_init, _ := http.NewRequest("PUT", "/vendor_users", bytes.NewBuffer(u_init_json))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req_init)

	u := VendorUser{
		UUID:  "",
		Email: "test3@localhost",
	}
	_ = json.Unmarshal(w.Body.Bytes(), &u.UUID)
	u_json, _ := json.Marshal(u)
	fmt.Println(string(u_json))
	req, _ := http.NewRequest("PATCH", "/vendor_users", bytes.NewBuffer(u_json))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
	req2, _ := http.NewRequest("GET", "/vendor_users", nil)
	r.ServeHTTP(w, req2)
	var u2 VendorUser
	_ = json.Unmarshal(w.Body.Bytes(), &u2)
	assert.Equal(t, "test3@localhost", u2.Email)
}
