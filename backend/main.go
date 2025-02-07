package main

import (
	"backend/containerHandler"
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

func main() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("databaseURL = ", databaseURL)

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
	defer DB.Close()

	// Создаем хэндлеры и передаем им ссылку на базу данных
	handler := &containerHandler.Handler{DB: DB}

	r := mux.NewRouter()
	r.HandleFunc("/containers", handler.GetContainers).Methods("GET")
	r.HandleFunc("/containers", handler.AddContainer).Methods("POST")

	// Разрешим политику CORS только для нашего фронтенда, http://localhost:3000 - для теста на локальном пк
	// TODO поменяй http://localhost:3000 на URL сервиса frontend поднятого в docker
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	// Оборачиваем наш маршрутизатор с CORS middleware
	http.Handle("/", handlers.CORS(originsOk, headersOk, methodsOk)(r))

	fmt.Println("Backend service started on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
