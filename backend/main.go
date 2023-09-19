package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"time"

	go_ora "github.com/sijms/go-ora/v2"
)

func main() {
	username := flag.String("user", "", "Gatorlink user")
	password := flag.String("password", "", "Oracle password")
	flag.Parse()
	if username == nil || *username == "" {
		panic("Username should not be empty")
	}
	if password == nil || *password == "" {
		panic("Password should not be empty")
	}
	port := 1521
	connStr := go_ora.BuildUrl("oracle.cise.ufl.edu", port, "orcl", *username, *password, nil)
	conn, err := sql.Open("oracle", connStr)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		cancel()
	}()
	err = conn.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<html><body>Hi, this is an API backend so you shouldn't be here")
	})
	http.HandleFunc("/api/v0/add-value", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer func() {
			cancel()
		}()

		if r.Method != http.MethodPost {
			writeError(w, "This requires POST")
			return
		}
		var args map[string]any

		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			writeError(w, fmt.Sprintln(err))
			return
		}
		valToInsert, ok := args["val"].(float64)
		if !ok {
			writeError(w, fmt.Sprintln(args["val"], "is not an integer but a", reflect.TypeOf(args["val"])))
			return
		}
		stmt, err := conn.PrepareContext(ctx, "INSERT INTO a(b) values(:1)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.ExecContext(ctx, valToInsert)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/api/v0/get-value", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer func() {
			cancel()
		}()

		if r.Method != http.MethodGet {
			writeError(w, "This requires get")
			return
		}
		stmt, err := conn.PrepareContext(ctx, "SELECT * FROM a")
		if err != nil {
			panic(err)
		}
		res, err := stmt.QueryContext(ctx)
		if err != nil {
			panic(err)
		}
		ints := []int{}
		for res.Next() {
			var v int
			err := res.Scan(&v)
			if err != nil {
				panic(err)
			}
			ints = append(ints, v)
		}
		msg := map[string]any{
			"result": ints,
		}
		b, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, string(b))
	})
	fmt.Println("Ready to serve")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func writeError(w http.ResponseWriter, m string) {
	msg := map[string]string{
		"error": m,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(b))

}
