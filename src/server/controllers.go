package main

import (
	"encoding/json" // Handling with JSON data
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

// Storing mock data

var data []Guitar

// Route handler
// Implementation of CRUD operations

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode("Server is running ")
}

func getAllSong(res http.ResponseWriter, req *http.Request) {
	var result []Guitar
	res.Header().Set("Content-type", "application/json")
	songdetails, err := db.Query("SELECT * from songdetails INNER JOIN song USING(songname) where songdetails.songname=song.songname;")
	if err != nil {
		panic(err.Error())
	}
	for songdetails.Next() {
		var song Song
		var details Guitar

		err = songdetails.Scan(&song.Name, &song.Tutorial, &song.Tab, &details.Singer, &details.URL)
		details.SongDetail = &song
		if err != nil {
			panic(err.Error())
		}
		result = append(result, details)
	}
	json.NewEncoder(res).Encode(result)
}

func getSong(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var song Guitar
	var songdetails Song
	// Getting the params from the request
	query := mux.Vars(req)
	singer := query["song"]
	result := db.QueryRow("SELECT * from song INNER JOIN songdetails USING(songname) WHERE song.songname=songdetails.songname and song.songname=?;", singer)
	err = result.Scan(&songdetails.Name, &songdetails.Tutorial, &songdetails.Tab, &song.Singer, &song.URL)
	if err != nil {
		panic(err.Error())
	}

	song.SongDetail = &songdetails
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(res).Encode(song)
}

func newEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newEntry Guitar
	json.NewDecoder(req.Body).Decode(&newEntry)
	_, err := db.Exec("INSERT INTO songdetails VALUES(?,?,?);", newEntry.SongDetail.Name, newEntry.SongDetail.Tutorial, newEntry.SongDetail.Tab)
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("INSERT INTO song VALUES(?,?,?);", newEntry.SongDetail.Name, newEntry.Singer, newEntry.URL)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(res).Encode(newEntry)

}

func updateEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var song Guitar

	json.NewDecoder(req.Body).Decode(&song)
	value := reflect.ValueOf(song)
	query := "UPDATE songdetails SET "
	for i := 0; i < value.NumField(); i++ {
		// fmt.Println(" ", query, value.Type().Field(i).Name, value.Field(i).Interface())
		// fmt.Println(query + value.Field(i).Interface())
	}
	// _,err:=db.Exec("UPDATE songdetails SET tutorial=?")
	// json.NewEncoder(res).Encode(updateEntry)
}

func deleteEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	query := mux.Vars(req)
	for index, item := range data {
		if item.Singer == query["song"] {
			data = append(data[:index], data[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(data)

}
