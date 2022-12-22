package config

import (
	"log"
	"os"
)

func GetPort() string {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port

}
