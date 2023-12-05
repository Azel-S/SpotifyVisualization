package database

import (
	"backend/utils"
	"net/http"
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

		totalTuples := 0
		// Execute query
		rows, err := db.database.Query(`
		SELECT T+S+R+Q+P+O+N FROM (SELECT COUNT(*) AS T FROM "SHAH.S".tracks), (SELECT COUNT(*) AS S FROM "SHAH.S".track_to_countries), (SELECT COUNT(*) AS R FROM "SHAH.S".artists),(SELECT COUNT(*) AS Q FROM "SHAH.S".artist_to_tracks), (SELECT COUNT(*) AS P FROM "SHAH.S".artist_to_genres), (SELECT COUNT(*) AS O FROM "SHAH.S".countries), (SELECT COUNT(*) AS N FROM "SHAH.S".country_to_code)
		`)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&totalTuples)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

		}

		utils.RespondWithJSON(w, http.StatusOK, totalTuples)
	}
}
