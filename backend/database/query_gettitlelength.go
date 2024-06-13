package database

// import (
// 	"backend/utils"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// // Gets popularity (From start_year to end_year)
// func (db *DB) GetTitleLength(w http.ResponseWriter, r *http.Request) {
// 	// Handles CORS and OPTIONS
// 	if !utils.HandleCORS(w, r) {
// 		// Only allow Get Methods
// 		if r.Method != http.MethodGet {
// 			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
// 			return
// 		}

// 		// Input structure
// 		var input struct {
// 			StartYear int
// 			EndYear   int
// 			Region_1  string
// 			Region_2  string
// 		}

// 		// Output structure
// 		var output struct {
// 			Years   []int     `json:"years"`
// 			Title_1 []float64 `json:"title_1"`
// 			Title_2 []float64 `json:"title_2"`
// 		}

// 		// Grab input values from url
// 		if r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" && r.URL.Query().Get("region_1") != "" && r.URL.Query().Get("region_2") != "" {
// 			var err error
// 			input.StartYear, err = strconv.Atoi(r.URL.Query().Get("start_year"))
// 			if err != nil {
// 				panic(err)
// 			}
// 			input.EndYear, err = strconv.Atoi(r.URL.Query().Get("end_year"))
// 			if err != nil {
// 				panic(err)
// 			}
// 			input.Region_1 = r.URL.Query().Get("region_1")
// 			input.Region_2 = r.URL.Query().Get("region_2")
// 			// region
// 		} else {
// 			utils.RespondWithError(w, http.StatusBadRequest, "start_year or end_year not specified")
// 			return
// 		}

// 		// Execute query
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
// 		defer cancel()
// 		rows, err := db.database.QueryContext(ctx,
// 			`
// 		WITH
// 		title_1 AS
// 		(
// 		        SELECT      t.release_year, c.region, MEDIAN(length(t.title)) title_1
// 				FROM        "SHAH.S".tracks t
// 				JOIN        track_to_countries ttc ON t.track_id = ttc.track_id
// 				JOIN        "SHAH.S".country_to_code ctc ON ttc.code = ctc.code
// 				JOIN        "SHAH.S".countries c ON ctc.country = c.name
// 				WHERE       c.region = :1 AND
// 		                    t.release_year >= :2 AND
// 		                    t.release_year <= :3
// 				GROUP BY    t.release_year, c.region
// 				ORDER BY    t.release_year
// 		),
// 		title_2 AS
// 		(
// 		        SELECT      t.release_year, c.region, MEDIAN(length(t.title)) title_2
// 				FROM        "SHAH.S".tracks t
// 				JOIN        track_to_countries ttc ON t.track_id = ttc.track_id
// 				JOIN        "SHAH.S".country_to_code ctc ON ttc.code = ctc.code
// 				JOIN        "SHAH.S".countries c ON ctc.country = c.name
// 				WHERE       c.region = :4 AND
// 		                    t.release_year >= :5 AND
// 		                    t.release_year <= :6
// 				GROUP BY    t.release_year, c.region
// 				ORDER BY    t.release_year
// 		)

// 		SELECT      t1.release_year, title_1, title_2
// 		FROM        title_1 t1
// 		JOIN        title_2 t2 ON t1.release_year = t2.release_year
// 		`, input.Region_1, input.StartYear, input.EndYear, input.Region_2, input.StartYear, input.EndYear)
// 		if err != nil {
// 			fmt.Println(err)
// 			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
// 			return
// 		}

// 		// Put result of query into output structure
// 		defer rows.Close()
// 		var (
// 			year    int
// 			title_1 float64
// 			title_2 float64
// 		)
// 		for rows.Next() {
// 			// Each row's values are put in temporary variables
// 			err = rows.Scan(&year, &title_1, &title_2)
// 			if err != nil {
// 				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
// 				return
// 			}

// 			// The temporary variables are appended to the output structure
// 			output.Years = append(output.Years, year)
// 			output.Title_1 = append(output.Title_1, title_1)
// 			output.Title_2 = append(output.Title_2, title_2)
// 		}

// 		utils.RespondWithJSON(w, http.StatusOK, output)
// 	}
// }
