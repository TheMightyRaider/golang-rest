package main

import (
	"encoding/json" // Handling with JSON data
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

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
	err = result.Scan(&songdetails.Name, &song.Singer, &song.URL, &songdetails.Tutorial, &songdetails.Tab)
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
	params := mux.Vars(req)

	type songdetails struct {
		Songname string
		Tab      string
		Tutorial string
		Url      string
		Singer   string
	}

	var song songdetails

	queryOne := "UPDATE songdetails SET "
	queryTwo := "UPDATE song SET "

	json.NewDecoder(req.Body).Decode(&song)

	value := reflect.ValueOf(song)
	for i := 0; i < value.NumField(); i++ {
		val := value.Field(i).Interface().(string)

		if value.Type().Field(i).Name == "Songname" && len(val) != 0 {
			queryOne += value.Type().Field(i).Name + "= '" + val + "'"
			queryTwo += value.Type().Field(i).Name + "= '" + val + "'"
		}

		if (value.Type().Field(i).Name == "Tutorial" || value.Type().Field(i).Name == "Tab") && len(val) != 0 {
			queryOne += ", " + value.Type().Field(i).Name + "= '" + val + "'"
		}

		if (value.Type().Field(i).Name == "Url" || value.Type().Field(i).Name == "Singer") && len(val) != 0 {
			queryTwo += ", " + value.Type().Field(i).Name + "= '" + val + "'"
		}

	}
	queryOne += " WHERE songname= '" + params["song"] + "';"
	queryTwo += " WHERE songname= '" + params["song"] + "';"

	_, err := db.Exec(queryOne)
	if err != nil {
		json.NewEncoder(res).Encode("Error")
	}
	_, err = db.Exec(queryTwo)
	if err != nil {
		json.NewEncoder(res).Encode("Error")
	}

	json.NewEncoder(res).Encode(song)
}

func deleteEntry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	// query := "DELETE songdetails, song FROM songdetails INNER JOIN song ON songdetails.songname=song.songname where songdetails.songname='" + params["song"] + "';"
	_, err = db.Exec("DELETE FROM song where songname=?;", params["song"])

	if err != nil {
		panic(err.Error())
	}

	_, err := db.Exec("DELETE FROM songdetails where songname=?;", params["song"])

	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(res).Encode("Deleted")

}
