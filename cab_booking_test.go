package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNearByCabs(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/booking/location?loc=narhe", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestBookCab(t *testing.T) {

	book := bookings{UserID: 1, CabID: 1, FromLoc: "katraj", ToLoc: "wakad"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(book)
	if errr != nil {
		log.Fatal(errr)
	}
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/booking/book", requestBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestPastBookings(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/booking/getBookings/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
