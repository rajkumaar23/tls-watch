package main

import (
	"log"
	"net/http"
	"tls-watch/api"
	"tls-watch/api/store"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load the env vars: %v", err)
	}
	store.InitializeDB()
}

func main() {
	auth, err := api.NewOIDCAuthenticator()
	if err != nil {
		log.Fatalf("failed to initialize the authenticator: %v", err)
	}

	router := api.NewRouter(auth)

	log.Print("server listening on http://localhost:2610/")
	if err := http.ListenAndServe("0.0.0.0:2610", router); err != nil {
		log.Fatalf("there was an error with the http server: %v", err)
	}
}
