package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/home", index).Methods("GET")
	return router
}

func TestIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/home", nil)
	response := httptest.NewRecorder()
	var body string
	Router().ServeHTTP(response, request)
	json.NewDecoder(response.Body).Decode(&body)
	assert.Equal(t, "Server is running", body, "Expected response")
}
