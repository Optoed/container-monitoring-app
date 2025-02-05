package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var db *sqlx.DB

func connectDB() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
}

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error when uploading .env file", err)
	}

	db, err = connectDB()
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	} else {
		log.Println("Successful connection to the database")
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/containers", GetContainers).Methods("GET")
	r.HandleFunc("/containers", AddContainer).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Backend service started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
