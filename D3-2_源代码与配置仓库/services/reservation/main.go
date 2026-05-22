package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/reservation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"service": "Reservation API",
			"status":  "Running ✨",
			"version": "v1.0.0",
		})
	})

	log.Println("Reservation service starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}