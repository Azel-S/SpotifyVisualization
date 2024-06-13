package database

import (
	"backend/utils"
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Gets followers of genre_1 and genre_2 (From start_year to end_year)
func (db *DB) CountTuples(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		collections := []*mongo.Collection{
			db.artist_to_genres,
			db.artist_to_tracks,
			db.artists,
			db.countries,
			db.country_to_code,
			db.track_to_countries,
			db.tracks,
		}

		count_pipeline := mongo.Pipeline{
			{{Key: "$count", Value: "tuples"}},
		}

		totalTuples := 0

		for _, collection := range collections {
			cursor, err := collection.Aggregate(context.TODO(), count_pipeline)
			if err != nil {
				fmt.Println("> ERROR: Could not count tuples for: ", collection.Name(), "%w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("Could not count tuples for: " + collection.Name() + err.Error()))
				return
			}

			for cursor.Next(context.TODO()) {
				var temp struct {
					Tuples int `bson:"tuples"`
				}

				err = cursor.Decode(&temp)
				if err != nil {
					fmt.Println("> ERROR: Could not decode tuples for: ", collection.Name(), "%w", err)
					utils.RespondWithError(w, http.StatusInternalServerError, ("Could not decode tuples for: " + collection.Name() + err.Error()))
					return
				}

				totalTuples += temp.Tuples
			}
		}

		utils.RespondWithJSON(w, http.StatusOK, totalTuples)
	}
}
