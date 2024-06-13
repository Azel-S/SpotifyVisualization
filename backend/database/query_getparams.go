package database

import (
	"backend/utils"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DB) GetYearRange(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		var output struct {
			StartYear int `json:"start_year" bson:"start_year"`
			EndYear   int `json:"end_year" bson:"end_year"`
		}

		// Get min_year
		min_pipeline := mongo.Pipeline{
			{{Key: "$group", Value: bson.D{{Key: "_id", Value: ""}, {Key: "start_year", Value: bson.D{{Key: "$min", Value: "$RELEASE_YEAR"}}}}}},
		}

		cursor, err := db.tracks.Aggregate(context.TODO(), min_pipeline)
		if err != nil {
			fmt.Println("> ERROR: Could not get start_year: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("Could not get start_year: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			err = cursor.Decode(&output)
			if err != nil {
				fmt.Println("> ERROR: Could not decode start_year: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("Could not decode start_year: " + err.Error()))
				return
			}
		}

		// Get max_year
		max_pipeline := mongo.Pipeline{
			{{Key: "$group", Value: bson.D{{Key: "_id", Value: ""}, {Key: "end_year", Value: bson.D{{Key: "$max", Value: "$RELEASE_YEAR"}}}}}},
		}

		cursor, err = db.tracks.Aggregate(context.TODO(), max_pipeline)
		if err != nil {
			fmt.Println("> ERROR: Could not get end_year: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("Could not get end_year: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			err = cursor.Decode(&output)
			if err != nil {
				fmt.Println("> ERROR: Could not decode end_year: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("Could not decode end_year: " + err.Error()))
				return
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}

func (db *DB) GetRegions(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		var output struct {
			Regions []string `json:"regions"`
		}

		regions_pipeline := mongo.Pipeline{
			{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$REGION"}}}}, {{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}}, {{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "REGION", Value: "$_id"}}}},
		}

		cursor, err := db.countries.Aggregate(context.TODO(), regions_pipeline)
		if err != nil {
			fmt.Println("> ERROR: Could not get regions: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("could not get regions: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			var temp struct {
				Region string `bson:"region"`
			}

			err = cursor.Decode(&temp)
			if err != nil {
				fmt.Println("> ERROR: Could not decode region: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("could not decode region: " + err.Error()))
				return
			}

			output.Regions = append(output.Regions, temp.Region)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}

func (db *DB) GetSubregions(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		var output struct {
			Subregions []string `json:"subregions"`
		}

		subregions_pipeline := mongo.Pipeline{
			{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$SUBREGION"}}}}, {{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}}, {{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "SUBREGION", Value: "$_id"}}}},
		}

		cursor, err := db.countries.Aggregate(context.TODO(), subregions_pipeline)
		if err != nil {
			fmt.Println("> ERROR: Could not get subregions: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("could not get subregions: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			var temp struct {
				Subregion string `bson:"subregion"`
			}

			err = cursor.Decode(&temp)
			if err != nil {
				fmt.Println("> ERROR: Could not decode subregions: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("could not decode subregions: " + err.Error()))
				return
			}

			output.Subregions = append(output.Subregions, temp.Subregion)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}

func (db *DB) GetGenres(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		var output struct {
			Genres []string `json:"genres"`
		}

		regions_pipeline := mongo.Pipeline{
			{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$GENRE"}}}}, {{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}}, {{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "GENRE", Value: "$_id"}}}},
		}

		cursor, err := db.artist_to_genres.Aggregate(context.TODO(), regions_pipeline)
		if err != nil {
			fmt.Println("> ERROR: Could not get genres: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("could not get genres: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			var temp struct {
				Genre string `bson:"genre"`
			}

			err = cursor.Decode(&temp)
			if err != nil {
				fmt.Println("> ERROR: Could not decode genres: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("could not decode genres: " + err.Error()))
				return
			}

			output.Genres = append(output.Genres, temp.Genre)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
