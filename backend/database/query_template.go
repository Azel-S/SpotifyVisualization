package database

import (
	"backend/utils"
	"net/http"
)

// Template Function, modify as needed
func (db *DB) Template(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		// Input Structure
		var input struct {
			// MyVariable type `json:"json_name"
			StartYear int `json:"start_year"`
		}

		// Output structure
		var output struct {
			// MyVariable type `json:"json_name"
			Year     []int     `json:"year"`
			Loudness []float64 `json:"loudness"`
		}

		// Decode given JSON into input structure
		err := utils.DecodeJSON(w, r, &input)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Execute query
		// :1 is the first variable passed in (StartYear in this case)
		rows, err := db.database.Query(`
			SELECT      t.release_year, MAX(t.loudness)
			FROM        "SHAH.S".tracks t
			WHERE       t.release_year >= :1
			GROUP BY    t.release_year
			ORDER BY    t.release_year ASC
			`, input.StartYear)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var (
			year     int
			loudness float64
		)
		for rows.Next() {
			// Each row's values are put in temporary variables
			err = rows.Scan(&year, &loudness)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			// The temporary variables are appended to the output structure
			output.Year = append(output.Year, year)
			output.Loudness = append(output.Loudness, loudness)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
