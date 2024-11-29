package main

// rename: lower-casing json variables to maintain json naming
type destination struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Description string   `json:"description"`
	Activities  []string `json:"activities"`
	BestSeason  string   `json:"best_season"`
	ImageURL    string   `json:"image_url"`
	Rating      float32  `json:"rating"`
	ReviewCount int      `json:"review_count"`
}

type review struct {
	ID            string `json:"id"`
	DestinationID string `json:"destination_id"`
	User          string `json:"user"`
	Rating        int    `json:"rating"`
	Comment       string `json:"comment"`
	Date          string `json:"date"`
}

var destinations = []destination{}
var reviews = []review{}
