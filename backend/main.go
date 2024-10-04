package main

import (
	"backend/database"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// import (
// 	"backend/database"
// 	"net/http"

// 	"github.com/joho/godotenv"
// )

// Get information from .env file (next to main.go file)
func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to retieve "+key+".env: %w", err))
	}

	return os.Getenv(key)
}

// func main() {
// 	// Connect to database (make sure UF VPN is active)
// 	db := database.DB{}
// 	db.Initalize(getEnv("USER"), getEnv("PASSWORD"))

// 	http.HandleFunc("/api/v0/CountTuples", db.CountTuples)

// 	// Set API links to be handled
// 	http.HandleFunc("/api/v0/GetYearRange", db.GetYearRange)
// 	http.HandleFunc("/api/v0/GetRegions", db.GetRegions)
// 	http.HandleFunc("/api/v0/GetSubregions", db.GetSubregions)
// 	http.HandleFunc("/api/v0/GetGenres", db.GetGenres)

// 	// Set API query links
// 	http.HandleFunc("/api/v0/GetPopularity", db.GetPopularity)
// 	http.HandleFunc("/api/v0/GetExplicit", db.GetExplicit)
// 	http.HandleFunc("/api/v0/GetGenrePopularity", db.GetGenrePopularity)
// 	http.HandleFunc("/api/v0/GetTitleLength", db.GetTitleLength)
// 	http.HandleFunc("/api/v0/GetAttributeComparison", db.GetAttributeComparison)

// 	// Start server
// 	fmt.Println("Serving at port: 8080")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	// Connect to database
	db := database.DB{}
	db.Initalize(getEnv("CONN_STR"))

	// Handle API requests
	http.HandleFunc("/api/v0/CountTuples", db.CountTuples)
	http.HandleFunc("/api/v0/GetYearRange", db.GetYearRange)
	http.HandleFunc("/api/v0/GetRegions", db.GetRegions)
	http.HandleFunc("/api/v0/GetSubregions", db.GetSubregions)
	http.HandleFunc("/api/v0/GetGenres", db.GetGenres)

	// Set API query links
	http.HandleFunc("/api/v0/GetPopularity", db.GetPopularity)
	// http.HandleFunc("/api/v0/GetExplicit", db.GetExplicit)
	// http.HandleFunc("/api/v0/GetGenrePopularity", db.GetGenrePopularity)
	// http.HandleFunc("/api/v0/GetTitleLength", db.GetTitleLength)
	// http.HandleFunc("/api/v0/GetAttributeComparison", db.GetAttributeComparison)

	// Start server
	fmt.Println("> Starting Sever (Port 8080)")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(fmt.Errorf("ERROR: Could not start server: %w", err))
	}
}
