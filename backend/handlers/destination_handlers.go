package handlers

import (
	"crypto/des"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tuoreste/API-Travel-Destinations.git/models"
)

//handling requests to fetch destinations by continent
func	GetDestinationsByContinent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	continent := vars["continent"]

	destinations := models.GetDestinationsByContinent(continent)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(destinations)
}

//get all continents
func	GetAllContinents(w http.ResponseWriter, r *http.Request) {
	continents := []string{
		"North America",
		"South America",
		"Europe",
		"Asia",
		"Africa",
		"Australia",
		"Antarctica",
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(continents)
}

//handling requests to fetch destinations by id
func	GetDestinationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	destination, found := models.GetDestinationByID(id)
	if !found {
		http.Error(w, "Destination not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(destination)
}
