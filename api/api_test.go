package api

import (
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
