package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tuoreste/API-Travel-Destinations.git/handlers"
)

func	main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8090"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})

	// the routes
	r.HandleFunc("/api/destinations/{continent}", handlers.GetDestinationsByContinent).Methods("GET")
	r.HandleFunc("/api/destinations/{id}", handlers.GetDestinationByID).Methods("GET")
	r.HandleFunc("/api/continents", handlers.GetAllContinents).Methods("GET")
	r.HandleFunc("/api/nearby", handlers.GetNearbyDestinations).Methods("GET")
	r.HandleFunc("/api/ws/geo-tracking", handlers.GeoTrackingWebSocket)

	http.Handle("/", r)

	handler := c.Handler(r)

	log.Println("Server starting on :8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
