package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func findPopularity2(w http.ResponseWriter, r *http.Request) {
	sql := "SELECT popularity, danceability, energy, Track_Genre.genre FROM Track JOIN Track_Genre ON Track_Genre.track_id=Track.track_id WHERE title like '%' || :1 ||'%' "
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer func() {
		cancel()
	}()

	if r.Method != http.MethodGet {
		writeError(w, "This requires get")
		return
	}
	stmt, err := conn.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	if t, ok := r.URL.Query()["title"]; len(t) == 1 && ok {
		res, err := stmt.QueryContext(ctx, t[0])
		if err != nil {
			panic(err)
		}
		popularityReturn2 := []PopularityReturn2{}
		for res.Next() {
			var popularity int
			var danceability float64
			var energy float64
			var genre string
			err := res.Scan(&popularity, &danceability, &energy, &genre)
			if err != nil {
				panic(err)
			}
			popularityReturn2 = append(popularityReturn2, PopularityReturn2{
				Popularity:   popularity,
				Danceability: danceability,
				Energy:       energy,
				Genre:        genre,
			})

		}
		msg := map[string]any{
			"result": popularityReturn2,
		}
		b, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, string(b))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, "There must be a title query")
	}
}

type PopularityReturn2 struct {
	Popularity   int     `json:"popularity"`
	Danceability float64 `json:"danceability"`
	Energy       float64 `json:"energy"`
	Genre        string  `json:"genre"`
}
