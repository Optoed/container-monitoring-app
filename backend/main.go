package main

import (
	"backend/containerHandler"
	"backend/repository"
	"backend/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := repository.InitDB()
	defer db.Close()
	if err != nil {
		log.Fatalln("error while initialization and connection to the db: ", err)
	}

	repo := &repository.Repository{DB: db}
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
	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
