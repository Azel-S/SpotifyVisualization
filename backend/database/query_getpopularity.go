package database

import (
	"backend/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// import (
// 	"backend/utils"
// 	"net/http"
// 	"strconv"
// )

// // Gets popularity (From start_year to end_year)
// func (db *DB) GetPopularity(w http.ResponseWriter, r *http.Request) {
// 	// Handles CORS and OPTIONS
// 	if !utils.HandleCORS(w, r) {
// 		// Only allow Get Methods
// 		if r.Method != http.MethodGet {
// 			utils.RespondWithError(w, http.StatusBadRequest, "GET method required")
// 			return
// 		}

// 		// Input structure
// 		var input struct {
// 			Attribute string
// 			StartYear int
// 			EndYear   int
// 		}

// 		// Output structure
// 		var output struct {
// 			Years        []int     `json:"years"`
// 			Popularities []float64 `json:"popularities"`
// 		}

// 		// Grab input values from url
// 		if r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" {
// 			input.Attribute = r.URL.Query().Get("attribute")
// 			input.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
// 			input.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
// 		} else {
// 			utils.RespondWithError(w, http.StatusBadRequest, "start_year or end_year not specified")
// 			return
// 		}

// 		var myQuery = `
// 		select t.release_year, (median(t.popularity)/(median(t.` + input.Attribute + `)+0.001))`

// 		myQuery = myQuery +
// 			`
// 			FROM "SHAH.S".TRACKS t
// 			where release_year >= :1 AND release_year <= :2
// 			GROUP BY RELEASE_YEAR
// 			ORDER BY RELEASE_YEAR
// 		`

// 		// Execute query
// 		rows, err := db.database.Query(myQuery, input.StartYear, input.EndYear)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
// 			return
// 		}

// 		// // Execute query
// 		// rows, err := db.database.Query(`
// 		// SELECT RELEASE_YEAR, (MEDIAN(popularity)/MEDIAN(loudness))
// 		// FROM "SHAH.S".TRACKS where release_year >= :1 AND release_year <= :2
// 		// GROUP BY RELEASE_YEAR
// 		// ORDER BY RELEASE_YEAR
// 		// 	`, input.StartYear, input.EndYear)
// 		// if err != nil {
// 		// 	utils.RespondWithError(w, http.StatusInternalServerError, ("Query exection failed: " + err.Error()))
// 		// 	return
// 		// }

// 		// Put result of query into output structure
// 		defer rows.Close()
// 		var (
// 			year       int
// 			popularity float64
// 		)
// 		for rows.Next() {
// 			// Each row's values are put in temporary variables
// 			err = rows.Scan(&year, &popularity)
// 			if err != nil {
// 				utils.RespondWithError(w, http.StatusInternalServerError, ("Row scan failed: " + err.Error()))
// 				return
// 			}

// 			// The temporary variables are appended to the output structure
// 			output.Years = append(output.Years, year)
// 			output.Popularities = append(output.Popularities, popularity)
// 		}

// 		utils.RespondWithJSON(w, http.StatusOK, output)
// 	}
// }

// Gets followers of genre_1 and genre_2 (From start_year to end_year)
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
			Attribute string
			StartYear int
			EndYear   int
		}

		// 		// Output structure
		var output struct {
			Years        []int     `json:"years"`
			Popularities []float64 `json:"popularities"`
		}

		// Grab input values from url
		if r.URL.Query().Get("attribute") != "" && r.URL.Query().Get("start_year") != "" && r.URL.Query().Get("end_year") != "" {
			input.Attribute = r.URL.Query().Get("attribute")
			input.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
			input.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "attribute, start_year or end_year not specified")
			return
		}

		group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$RELEASE_YEAR"}, {Key: "Popularity", Value: bson.D{{Key: "$median", Value: bson.D{{Key: "input", Value: "$POPULARITY"}, {Key: "method", Value: "approximate"}}}}}, {Key: "Attribute", Value: bson.D{{Key: "$median", Value: bson.D{{Key: "input", Value: "$" + strings.ToUpper(input.Attribute)}, {Key: "method", Value: "approximate"}}}}}}}}
		match := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: []interface{}{bson.D{{Key: "_id", Value: bson.D{{Key: "$gte", Value: input.StartYear}, {Key: "$lte", Value: input.EndYear}}}}}}}}}
		project := bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "YEAR", Value: "$_id"}, {Key: "Popularity", Value: 1}, {Key: "Attribute", Value: 1}}}}
		sort := bson.D{{Key: "$sort", Value: bson.D{{Key: "YEAR", Value: 1}}}}

		cursor, err := db.tracks.Aggregate(context.TODO(), mongo.Pipeline{group, match, project, sort})
		if err != nil {
			fmt.Println("> ERROR: Could not get popularity: %w", err)
			utils.RespondWithError(w, http.StatusInternalServerError, ("could not get popularity: " + err.Error()))
			return
		}

		for cursor.Next(context.TODO()) {
			var temp struct {
				Year       int     `bson:"year"`
				Popularity float64 `bson:"popularity"`
				Attribute  float64 `bson:"attribute"`
			}

			err = cursor.Decode(&temp)
			if err != nil {
				fmt.Println("> ERROR: Could not decode popularity: %w", err)
				utils.RespondWithError(w, http.StatusInternalServerError, ("could not decode popularity: " + err.Error()))
				return
			}

			output.Years = append(output.Years, temp.Year)
			output.Popularities = append(output.Popularities, (temp.Popularity / (temp.Attribute + 0.001)))
		}

		utils.RespondWithJSON(w, http.StatusOK, output)
	}
}
