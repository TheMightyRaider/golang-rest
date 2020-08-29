package main

import (
	"encoding/json" // Handling with JSON data
	"github.com/gorilla/mux"
	"net/http"
)

// Route handler
// Implementation of CRUD operations

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode("Server is running ")
}

func getAllSong(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func getSong(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	// Getting the params from the request
	query := mux.Vars(req)
	for _, item := range data {
		if item.Singer == query["singer"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func newEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newEntry Guitar
	json.NewDecoder(req.Body).Decode(&newEntry)
	data = append(data, newEntry)
	json.NewEncoder(res).Encode(newEntry)

}

func updateEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	query := mux.Vars(req)
	var updateEntry Guitar
	for index, item := range data {
		if item.Song.Name == query["song"] {
			json.NewDecoder(req.Body).Decode(&updateEntry)
			data[index] = updateEntry
			break
		}
	}
	json.NewEncoder(res).Encode(updateEntry)
}

func deleteEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	query := mux.Vars(req)
	for index, item := range data {
		if item.Song.Name == query["song"] {
			data = append(data[:index], data[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(data)

}
