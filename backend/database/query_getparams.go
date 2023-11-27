package database

import (
	"backend/utils"
	"net/http"
)

func (db *DB) GetYearRange(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		// Output structure
		var output struct {
			StartYear int `json:"start_year"`
			EndYear   int `json:"end_year"`
		}

		// Execute query
		rows1, err := db.database.Query(`
		SELECT      MIN(release_year)
		FROM        "SHAH.S".tracks
		`)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows1.Close()
		for rows1.Next() {
			err = rows1.Scan(&output.StartYear)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}
		}

		// Execute query
		rows2, err := db.database.Query(`
		SELECT      MAX(release_year)
		FROM        "SHAH.S".tracks
		`)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows2.Close()
		for rows2.Next() {
			err = rows2.Scan(&output.EndYear)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
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

		// Output structure
		var output struct {
			Regions []string `json:"regions"`
		}

		// Execute query
		rows, err := db.database.Query(`
		SELECT      distinct region
		FROM        "SHAH.S".countries
		ORDER BY	region ASC
		`)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var region string
		for rows.Next() {
			// Each row's values are put in temporary variables
			err = rows.Scan(&region)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			// The temporary variables are appended to the output structure
			output.Regions = append(output.Regions, region)
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

		// Output structure
		var output struct {
			Genres []string `json:"genres"`
		}

		// Execute query
		rows, err := db.database.Query(`
		SELECT      distinct genre
		FROM        "SHAH.S".artist_to_genres
		ORDER BY	genre ASC
		`)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var genre string
		for rows.Next() {
			err = rows.Scan(&genre)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			output.Genres = append(output.Genres, genre)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
