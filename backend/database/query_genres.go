package database

import (
	"backend/utils"
	"net/http"
	"strconv"
)

// Gets followers of genre_1 and genre_2 (From start_year to end_year)
func (db *DB) GetGenreFollowers(w http.ResponseWriter, r *http.Request) {
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
			Genre_1   string
			Genre_2   string
		}

		// Output structure
		var output struct {
			Years       []int `json:"years"`
			Followers_1 []int `json:"followers_1"`
			Followers_2 []int `json:"followers_2"`
		}

		// Grab input values from url
		if r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" && r.URL.Query().Get("genre_1") != "" && r.URL.Query().Get("genre_2") != "" {
			input.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
			input.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
			input.Genre_1 = r.URL.Query().Get("genre_1")
			input.Genre_2 = r.URL.Query().Get("genre_2")
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "start_year, end_year, genre_1, or genre_2 not specified")
			return
		}

		// Execute query
		rows, err := db.database.Query(`
		WITH genre_1 AS
		(
		    SELECT      t.release_year, SUM(a.popularity) followers_1
		    FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
		    WHERE       t.track_id = att.track_id AND
		                att.artist_id = a.artist_id AND
		                a.artist_id = atg.artist_id AND
		                atg.genre = :1 AND
		                t.release_year >= :2 AND
		                t.release_year <= :3
		    GROUP BY    t.release_year
		    ),
		genre_2 AS
		(
		    SELECT      t.release_year, SUM(a.popularity) followers_2
		    FROM        "SHAH.S".tracks t, "SHAH.S".artist_to_tracks att, "SHAH.S".artists a, "SHAH.S".artist_to_genres atg
		    WHERE       t.track_id = att.track_id AND
		                att.artist_id = a.artist_id AND
		                a.artist_id = atg.artist_id AND
		                atg.genre = :4 AND
		                t.release_year >= :5 AND
		                t.release_year <= :6
		    GROUP BY    t.release_year
		)

		SELECT      *
		FROM        genre_1 NATURAL JOIN genre_2
		ORDER BY    release_year ASC
		`, input.Genre_1, input.StartYear, input.EndYear, input.Genre_2, input.StartYear, input.EndYear)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
			return
		}

		// Put result of query into output structure
		defer rows.Close()
		var (
			year        int
			followers_1 int
			followers_2 int
		)
		for rows.Next() {
			err = rows.Scan(&year, &followers_1, &followers_2)
			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
				return
			}

			output.Years = append(output.Years, year)
			output.Followers_1 = append(output.Followers_1, followers_1)
			output.Followers_2 = append(output.Followers_2, followers_2)
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
