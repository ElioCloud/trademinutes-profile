package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"trademinutes-profile/config"
	"trademinutes-profile/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println(".env file not found, assuming production environment variables")
		}
	}

	// Connect to MongoDB
	config.ConnectDB()
	fmt.Println("✅ Connected to MongoDB:", config.GetDB().Name())

	// Create router
	router := mux.NewRouter()

	// Health check (public)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// Register routes (prefix handled inside routes.ProfileRoutes)
	routes.ProfileRoutes(router)

	// Optional: log unmatched routes
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("404 Not Found: %s %s\n", r.Method, r.URL.Path)
		http.Error(w, "404 not found", http.StatusNotFound)
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Profile service running on :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
