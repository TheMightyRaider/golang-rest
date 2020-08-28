package main

import (
	"encoding/json"          // Handling with JSON data
	"github.com/gorilla/mux" // handler for http request aka router
	"log"                    // For logging errors
	"net/http"               // communicating with the server
)

// Guitar based playlist (Model)
type Guitar struct {
	Song   *Song  `json:"name"` // Reference a another model
	Singer string `json:"singer"`
	URL    string `json:"link"`
}

// Song provides the composing details
type Song struct {
	Name     string `json:"name"`
	Tutorial string `json:"chords"`
	Tab      string `json:"Tab"`
}

// Storing mock data

var data []Guitar

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

func main() {
	//Init Router
	r := mux.NewRouter()

	// Adding data
	data = append(data, Guitar{Song: &Song{Name: "Hurt", Tutorial: "https://www.youtube.com/watch?v=hEKexrWq5RU", Tab: "https://tabs.ultimate-guitar.com/tab/johnny-cash/hurt-chords-89849"}, Singer: "Johnny Cash", URL: "https://www.youtube.com/watch?v=8AHCfZTRGiI"})
	data = append(data, Guitar{Song: &Song{Name: "Viva la Vida", Tutorial: "https://www.youtube.com/watch?v=dO7t54nT90U", Tab: "https://tabs.ultimate-guitar.com/tab/coldplay/viva-la-vida-tabs-717746"}, Singer: "ColdPlay", URL: "https://www.youtube.com/watch?v=-ZvsGmYKhcU"})
	data = append(data, Guitar{Song: &Song{Name: "The Night We Met", Tutorial: "https://www.youtube.com/watch?v=kEcbaH45whw&t=150s", Tab: "https://tabs.ultimate-guitar.com/tab/lord-huron/the-night-we-met-tabs-1706742"}, Singer: "Lord Huron", URL: "https://www.youtube.com/watch?v=KtlgYxa6BMU"})
	data = append(data, Guitar{Song: &Song{Name: "Apocalypse ", Tutorial: "https://www.youtube.com/watch?v=3WL8CBGJYg4&t=20s", Tab: "https://tabs.ultimate-guitar.com/tab/cigarettes-after-sex/apocalypse-tabs-2354393"}, Singer: "Cigarettes After Sex", URL: "https://www.youtube.com/watch?v=sElE_BfQ67s"})

	//Router for handling http requests aka endpoints

	r.HandleFunc("/home", index).Methods("GET")
	r.HandleFunc("/song/all", getAllSong).Methods("GET")
	r.HandleFunc("/song/create", newEntry).Methods("POST")
	r.HandleFunc("/song/{singer}", getSong).Methods("GET")
	r.HandleFunc("/song/update/{song}", updateEntry).Methods("PUT")
	r.HandleFunc("/song/delete/{song}", deleteEntry).Methods("DELETE")

	// Initiating the Server

	log.Fatal(http.ListenAndServe(":8000", r))
}
