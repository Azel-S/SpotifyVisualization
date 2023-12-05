package database

import (
	"backend/utils"
	"net/http"
	"strconv"
)

// Gets popularity (From start_year to end_year)
func (db *DB) GetAttributeComparison(w http.ResponseWriter, r *http.Request) {
	// Handles CORS and OPTIONS
	if !utils.HandleCORS(w, r) {
		// Only allow Get Methods
		if r.Method != http.MethodGet {
			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
			return
		}

		// Input structure
		var input struct {
			Attribute1 string
			Attribute2 string
			Genre      string

			StartYear int
			EndYear   int
		}

		// Output structure
		var output struct {
			Years      []int     `json:"years"`
			Attribute1 []float64 `json:"attribute_1"`
			Attribute2 []float64 `json:"attribute_2"`
		}

		// Grab input values from url
		if r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" && r.URL.Query().Get("attribute_1") != "" && r.URL.Query().Get("attribute_2") != "" && r.URL.Query().Get("genre") != "" {
			input.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
			input.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
			input.Attribute1 = r.URL.Query().Get("attribute_1")
			input.Attribute2 = r.URL.Query().Get("attribute_2")
			input.Genre = r.URL.Query().Get("genre")
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "start_year or end_year or artist_name not specified")
			return
		}

		var myQuery = `
		select t.release_year, avg(t.` + input.Attribute1 + `), avg(t.` + input.Attribute2 + `)`

		myQuery = myQuery +
			`
			FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
			WHERE       t.track_id = att.track_id AND
		                att.artist_id = a.artist_id AND
		                a.artist_id = atg.artist_id AND
		                atg.genre = :1 AND
						t.release_year >= :2 AND
						t.release_year <= :3	
			GROUP BY    t.release_year
			ORDER BY    t.release_year ASC
		`

		// Execute query
		rows, err := db.database.Query(myQuery, input.Genre, input.StartYear, input.EndYear)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var (
			year        int
			attribute_1 float64
			attribute_2 float64
		)
		for rows.Next() {
			// Each row's values are put in temporary variables
			err = rows.Scan(&year, &attribute_1, &attribute_2)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			// The temporary variables are appended to the output structure
			output.Years = append(output.Years, year)
			output.Attribute1 = append(output.Attribute1, attribute_1)
			output.Attribute2 = append(output.Attribute2, attribute_2)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
