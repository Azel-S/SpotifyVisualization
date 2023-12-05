package main

import (
	"backend/database"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Get information from .env file (next to main.go file)
func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to retieve "+key+".env: %w", err))
	}

	return os.Getenv(key)
}

func main() {
	// Connect to database (make sure UF VPN is active)
	db := database.DB{}
	db.Initalize(getEnv("USER"), getEnv("PASSWORD"))

	http.HandleFunc("/api/v0/CountTuples", db.CountTuples)

	// Set API links to be handled
	http.HandleFunc("/api/v0/GetYearRange", db.GetYearRange)
	http.HandleFunc("/api/v0/GetRegions", db.GetRegions)
	http.HandleFunc("/api/v0/GetSubregions", db.GetSubregions)
	http.HandleFunc("/api/v0/GetGenres", db.GetGenres)

	// Set API query links
	http.HandleFunc("/api/v0/GetPopularity", db.GetPopularity)
	http.HandleFunc("/api/v0/GetExplicit", db.GetExplicit)
	http.HandleFunc("/api/v0/GetGenreFollowers", db.GetGenreFollowers)
	http.HandleFunc("/api/v0/GetAvgDuration", db.GetAvgDuration)
	http.HandleFunc("/api/v0/GetTitleLength", db.GetTitleLength)
	http.HandleFunc("/api/v0/GetAttributeComparison", db.GetAttributeComparison)
	// http.HandleFunc("/api/v0/query_1", db.query_2)

	// Start server
	fmt.Println("Serving at port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
