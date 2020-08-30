package main

// Guitar based playlist (Model)
type Guitar struct {
	SongDetail *Song  `json:"Details"` // Reference a another model
	Singer     string `json:"Singer"`
	URL        string `json:"Link"`
}

// Song provides the composing details
type Song struct {
	Name     string `json:"Name"`
	Tutorial string `json:"Tutorial"`
	Tab      string `json:"Tab"`
}
