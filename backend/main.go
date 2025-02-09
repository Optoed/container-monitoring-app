package main

import (
	"backend/containerHandler"
	"backend/repository"
	"backend/service"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var DB *sqlx.DB

func initDB() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//log.Println("databaseURL = ", databaseURL)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = sqlx.Connect("postgres", databaseURL)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	} else {
		log.Println("Successful connection to the database")
	}
}

func main() {
	initDB()
	defer DB.Close()

	repo := &repository.Repository{DB: DB}
	serv := &service.Service{Repo: repo}
	handler := &containerHandler.Handler{Service: serv}

	r := mux.NewRouter()
	r.HandleFunc("/containers", handler.GetContainers).Methods("GET")

	// Разрешим политику CORS только для собственного фронтенда,
	// http://localhost:3000 - для взаимодействия на локальном пк
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	containerizedFrontendURL := os.Getenv("FRONTEND_URL")
	localFrontendURL := "http://localhost:3000"
	originsOk := handlers.AllowedOrigins([]string{containerizedFrontendURL, localFrontendURL})

	// Оборачиваем наш маршрутизатор с CORS middleware
	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(r))

	//Start consumer
	go serv.StartConsume()

	log.Println("Attempting to start HTTP server...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
