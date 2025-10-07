package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Address string
	Clients []string
}

func main() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got new request from %s", r.RemoteAddr)

		for _, client := range config.Clients {
			if _, err := http.Post(client, "text/plain", nil); err != nil {
				log.Printf("Failed to access %v: %v", client, err)
			}
		}
	})

	log.Printf("Starting schedule relay on %s...", config.Address)
	http.ListenAndServe(config.Address, nil)
}
