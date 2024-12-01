package handlers

import (
	// "crypto/des"
	"encoding/json"
	// "go/constant"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	// "golang.org/x/net/websocket"

	// "github.com/pelletier/go-toml/query"
	"github.com/tuoreste/API-Travel-Destinations.git/models"
)

type	GeoLocationRequest struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius	 float64 `json:"radius"`
}


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

// calc dist btn 2 geo-coord
func	CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const	R = 6371
	latDiff := (lat2 - lat1) * (math.Pi / 180)
	lonDiff := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(latDiff/2)*math.Sin(latDiff/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
		math.Sin(lonDiff/2)*math.Sin(lonDiff/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// handle a req to find nearby dest
func	GetNearbyDestinations(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	lat, _ := strconv.ParseFloat(query.Get("latitude"), 64)
	lon, _ := strconv.ParseFloat(query.Get("longitude"), 64)
	radius, _ := strconv.ParseFloat(query.Get("radius"), 64)

	type NearbyDestination struct {
		models.Destination
		Notification string `json:"notification"`
	}

	var nearbyDestinations []NearbyDestination

	for _, dest := range models.Destinations {
		distance := CalculateDistance(lat, lon, dest.Location.Latitude, dest.Location.Longitude)
		if distance <= radius {
			notification := "You are near " + dest.Name + ". Check out the highlights: " + strings.Join(dest.Highlights, ", ")

			log.Printf("Notification: %s", notification)

			nearbyDestinations = append(nearbyDestinations, NearbyDestination{
				Destination:  dest,
				Notification: notification,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(nearbyDestinations)
}

var	upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handling websocket here on rea-time notifi
func	GeoTrackingWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		var locationReq	GeoLocationRequest
		err := conn.ReadJSON(&locationReq)
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		//call nearby locations
		var notifications []string
		for _, dest := range models.Destinations {
			distance := CalculateDistance(locationReq.Latitude, locationReq.Longitude, dest.Location.Latitude, dest.Location.Longitude)
			if distance <= locationReq.Radius {
				highlights := strings.Join(dest.Highlights, ", ")
				notification := "You are near " + dest.Name + ". Highlights: " + highlights
				notifications = append(notifications, notification)
			}
		}
		if err := conn.WriteJSON(notifications); err != nil {
			log.Println("WebSocket write error:", err)
			break
		}
	}
}
