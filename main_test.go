package main

import (
	"go-restapi-mysql/app/Controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHomepageHandlerSuccess(t *testing.T) {
	// mockResponse := `{"status":true}`
	r := SetUpRouter()
	r.GET("/", Controllers.GetIndex)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// responseData, _ := ioutil.ReadAll(w.Body)
	// assert.Equal(t, mockResponse, string(responseData))
	// assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestHomepageHandlerFail(t *testing.T) {
	// mockResponse := `{"status":true}`
	r := SetUpRouter()
	r.GET("/", Controllers.GetIndex)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
