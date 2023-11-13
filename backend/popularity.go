package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func findPopularity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer func() {
		cancel()
	}()

	if r.Method != http.MethodGet {
		writeError(w, "This requires get")
		return
	}
	sql := "SELECT EXTRACT(YEAR FROM release_date), sum(popularity) FROM Track GROUP BY EXTRACT(YEAR FROM release_date) ORDER BY EXTRACT(YEAR FROM release_date) DESC"
	stmt, err := conn.PrepareContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	res, err := stmt.QueryContext(ctx)
	if err != nil {
		panic(err)
	}

	yearArr := []int{}
	popularityArr := []int{}
	for res.Next() {
		var year int
		var popularity int
		err := res.Scan(&year, &popularity)
		if err != nil {
			panic(err)
		}
		yearArr = append(yearArr, year)
		popularityArr = append(popularityArr, popularity)
	}
	msg := map[string][]int{
		"years":      yearArr,
		"popularity": popularityArr,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
