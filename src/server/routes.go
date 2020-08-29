package main

import "github.com/gorilla/mux"

func route(r *mux.Router) {
	r.HandleFunc("/home", index).Methods("GET")
	r.HandleFunc("/song/all", getAllSong).Methods("GET")
	r.HandleFunc("/song/create", newEntry).Methods("POST")
	r.HandleFunc("/song/{singer}", getSong).Methods("GET")
	r.HandleFunc("/song/update/{song}", updateEntry).Methods("PUT")
	r.HandleFunc("/song/delete/{song}", deleteEntry).Methods("DELETE")
}
