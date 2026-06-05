package config

import (
	"log"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT is not set in environment variables, fallback to default port: %s", port)
	}
	return port
}
