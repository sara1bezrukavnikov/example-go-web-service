package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT not set. Using default %s.", port)
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "World"
		log.Printf("APP_NAME not set. Using default %s.", appName)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
		log.Printf("HOST not set. Using default %s.", host)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received for /")
		fmt.Fprintf(w, "Hello, %s! You are running on %s.\n", appName, host)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
