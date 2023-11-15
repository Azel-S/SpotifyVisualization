package database

import (
	"backend/utils"
	"net/http"
)

// Gets popularity (From start_year to end_year)
func (db *DB) GetPopularity(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		// Input Structure
		var input struct {
			StartYear int `json:"start_year"`
			EndYear   int `json:"end_year"`
		}

		// Output structure
		var output struct {
			Year       []int     `json:"year"`
			Popularity []float64 `json:"popularity"`
		}

		// Decode given JSON into input structure
		err := utils.DecodeJSON(w, r, &input)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Execute query
		rows, err := db.database.Query(`
			SELECT      t.release_year, MAX(t.popularity)
			FROM        "SHAH.S".tracks t
			WHERE       t.release_year >= :1 AND
						t.release_year <= :2
			GROUP BY    t.release_year
			ORDER BY    t.release_year ASC
			`, input.StartYear, input.EndYear)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var (
			year       int
			popularity float64
		)
		for rows.Next() {
			// Each row's values are put in temporary variables
			err = rows.Scan(&year, &popularity)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			// The temporary variables are appended to the output structure
			output.Year = append(output.Year, year)
			output.Popularity = append(output.Popularity, popularity)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
