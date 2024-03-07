package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type CorsError struct {
	err string
}

func (e *CorsError) Error() string {
	return e.err
}

func SetupCors(m *http.ServeMux) (http.Handler, error) {
	mode := os.Getenv("MODE")

	var (
		options        cors.Options
		corsMiddleware *cors.Cors
	)

	if mode == "DEV" {
		corsMiddleware = cors.Default()
		fmt.Println("DEV MODE")
	} else if mode == "PROD" {
		options = cors.Options{
			AllowedOrigins: []string{"https://naequina.com"},
			AllowedMethods: []string{"POST"},
			AllowedHeaders: []string{"*"},
		}

		corsMiddleware = cors.New(options)
	} else {
		return nil, &CorsError{err: "Invalid mode (" + mode + ")"}
	}

	return corsMiddleware.Handler(m), nil
}

func SetupCORSHttp(w *http.ResponseWriter, req *http.Request) {
	mode := os.Getenv("MODE")

	if mode == "DEV" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	} else if mode == "PROD" {
		(*w).Header().Set("Access-Control-Allow-Origin", "https://naequina.com, https://www.naequina.com")
		(*w).Header().Set("Access-Control-Allow-Methods", "POST")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	} else {
		log.Fatal("error seting up cors")
	}

}
