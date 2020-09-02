package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // For communicating with the database
	"github.com/gorilla/mux"           // handler for http request aka router
	"log"                              // For logging errors
	"net/http"                         // communicating with the server
)

var db, err = sql.Open("mysql", "mighty:Madhurima.k7@tcp(127.0.0.1:3306)/")

func main() {
	//Init Router
	r := mux.NewRouter()

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("USE GuitarMusic;")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connected to the Database")

	//Router for handling http requests aka endpoints
	route(r)

	// Initiating the Server
	log.Fatal(http.ListenAndServe(":8000", r))
}
