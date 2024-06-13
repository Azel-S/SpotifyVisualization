package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client   *mongo.Client
	database *mongo.Database

	// collections (tables)
	artist_to_genres   *mongo.Collection
	artist_to_tracks   *mongo.Collection
	artists            *mongo.Collection
	countries          *mongo.Collection
	country_to_code    *mongo.Collection
	track_to_countries *mongo.Collection
	tracks             *mongo.Collection
}

// Sets up connection to MongoDB database using conn_str (in .env file)
func (db *DB) Initalize(conn_str string) {
	opts := options.Client().ApplyURI(conn_str).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	var err error

	// Create a new client and connect to the server
	db.client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(fmt.Errorf("> ERROR: Could not connect to MongoDB: %w", err))
	}

	// Make sure connection is solid by pinging
	err = db.client.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Errorf("> ERROR: Could not ping to MongoDB: %w", err))
	}

	fmt.Println("> Connected to MongoDB!")

	// TODO: Error validation
	db.database = db.client.Database("spotify_trends_db")

	db.artist_to_genres = db.database.Collection("artist_to_genres")
	db.artist_to_tracks = db.database.Collection("artist_to_tracks")
	db.artists = db.database.Collection("artists")
	db.countries = db.database.Collection("countries")
	db.country_to_code = db.database.Collection("country_to_code")
	db.track_to_countries = db.database.Collection("track_to_countries")
	db.tracks = db.database.Collection("tracks")

}
