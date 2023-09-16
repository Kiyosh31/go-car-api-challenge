package main

import (
	"cars-go/api"
	"cars-go/config"
	"cars-go/store"
	"log"
)

func main() {
	env, err := config.LoadEnvVars()
	if err != nil {
		log.Fatalf("Could not load env vars: %v", err)
	}

	store := store.NewCarStore()
	server := api.NewServer(env.Port, *store)

	err = server.Start()
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}

	log.Printf("Server running in port: %v", env.Port)
}
