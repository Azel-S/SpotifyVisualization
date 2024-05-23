package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://" + getEnv("USER") + ":" + getEnv("PASSWORD") + "@cluster-0.ru4ftmg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster-0").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
