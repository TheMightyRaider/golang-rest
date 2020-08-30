package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func database() {

	_, err = db.Exec("USE GuitarMusic;")
	if err != nil {
		panic(err.Error())

	}

	_, err = db.Exec("INSERT INTO songdetails VALUES('Hurt','https://www.youtube.com/watch?v=hEKexrWq5RU','https://tabs.ultimate-guitar.com/tab/johnny-cash/hurt-chords-89849');")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO song VALUES('Hurt','Johnny Cash','https://www.youtube.com/watch?v=8AHCfZTRGiI');")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO songdetails VALUES('Viva la Vida','https://www.youtube.com/watch?v=dO7t54nT90U','https://tabs.ultimate-guitar.com/tab/coldplay/viva-la-vida-tabs-717746');")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO song VALUES('Viva la Vida','ColdPlay','https://www.youtube.com/watch?v=-ZvsGmYKhcU');")
	if err != nil {
		panic(err)
	}

	defer db.Close()

}
