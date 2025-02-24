package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

// create a test to check the status of the health endpoint
func TestHealth(t *testing.T) {
	// create a new router
	router := gin.Default()
	// add the health endpoint to the router
	router.GET("/health", getHealth)
	// create a new request to the health endpoint
	req, _ := http.NewRequest("GET", "/health", nil)
	// create a new response recorder
	w := httptest.NewRecorder()
	// execute the request on the router
	router.ServeHTTP(w, req)
	// check the status code
	assert.Equal(t, 200, w.Code)
	// check the response body
	require.JSONEq(t, `[{"status":"okay"}]`, w.Body.String())
}
