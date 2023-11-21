package database

import (
	"backend/utils"
	"net/http"
	"strconv"
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

		// Input structure
		var input struct {
			StartYear int
			EndYear   int
		}

		// Output structure
		var output struct {
			Years        []int     `json:"years"`
			Popularities []float64 `json:"popularities"`
		}

		// Grab input values from url
		if r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" {
			input.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
			input.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "start_year or end_year not specified")
			return
		}

		// Execute query
		rows, err := db.database.Query(`
			SELECT      t.release_year, SUM(t.popularity)
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
			output.Years = append(output.Years, year)
			output.Popularities = append(output.Popularities, popularity)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
