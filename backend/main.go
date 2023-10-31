package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	go_ora "github.com/sijms/go-ora/v2"
)

// Get information from .env file.
func getEnv(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var conn *sql.DB

func main() {
	username := flag.String("user", getEnv("DBUSER"), "Gatorlink user")
	password := flag.String("password", getEnv("PASSWORD"), "Oracle password")
	flag.Parse()
	if username == nil || *username == "" {
		panic("Username should not be empty")
	}
	if password == nil || *password == "" {
		panic("Password should not be empty")
	}
	port := 1521
	connStr := go_ora.BuildUrl("oracle.cise.ufl.edu", port, "orcl", *username, *password, nil)
	var err error
	conn, err = sql.Open("oracle", connStr)
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, "<html><body>Hi, this is an API backend so you shouldn't be here")
	})
	http.HandleFunc("/api/v0/add-value", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer func() {
			cancel()
		}()

		if r.Method != http.MethodPost {
			writeError(w, "This requires POST")
			return
		}
		var args map[string]any
		b, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(b, &args)
		if err != nil {
			writeError(w, fmt.Sprintln(err))
			return
		}
		fmt.Println(args)
		valToInsert, ok := args["message"].(float64)
		if !ok {
			writeError(w, fmt.Sprintln(args["message"], "is not an integer but a", reflect.TypeOf(args["message"])))
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
	http.HandleFunc("/api/v0/get-popularity", findPopularity)
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
