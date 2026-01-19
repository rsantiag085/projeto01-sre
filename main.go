package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Cargo struct {
	ID        string    `json:"id"`
	Item      string    `json:"item"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK - System Healthy")
}

func cargoHandler(w http.ResponseWriter, r *http.Request) {
	inventory := []Cargo{
		{ID: "1", Item: "Servidores Dell", Quantity: 5, Status: "Em Trânsito", Timestamp: time.Now()},
		{ID: "2", Item: "Switches Cisco", Quantity: 10, Status: "Armazenado", Timestamp: time.Now()},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func main() {
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/inventory", cargoHandler)

	port := ":8080"
	fmt.Printf("Servidor Go-Cargo rodando na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}