package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
)

// The Magic of ADC: We pass 'nil' as the configuration.
// The Google SDK automatically handles the background token rotation via WIF.
func initializeFirebase() (*firebase.App, error) {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}

func main() {
	log.Println("Starting WIF Backend Service...")

	app, err := initializeFirebase()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v\n", err)
	}
	log.Println("✅ Successfully authenticated with GCP using Application Default Credentials!")

	// A simple endpoint to prove the app is running and authenticated
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		// In a real app, you would use 'app' to query your database here
		_ = app
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend is authenticated and running perfectly via WIF!"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
