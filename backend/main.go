package main

import (
	"backend/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var DB *sqlx.DB

func connectDB() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
}

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error when uploading .env file", err)
	}

	DB, err = connectDB()
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	} else {
		log.Println("Successful connection to the database")
	}
	defer DB.Close()

	// Создаем хэндлеры и передаем им ссылку на базу данных
	handler := &handlers.Handler{DB: DB}

	r := mux.NewRouter()
	r.HandleFunc("/containers", handler.GetContainers).Methods("GET")
	r.HandleFunc("/containers", handler.AddContainer).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Backend service started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
