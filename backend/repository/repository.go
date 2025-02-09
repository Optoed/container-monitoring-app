package repository

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

type Repository struct {
	DB *sqlx.DB
}

func InitDB() (*sqlx.DB, error) {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//log.Println("databaseURL = ", databaseURL)

	var DB *sqlx.DB
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
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	} else {
		log.Println("Successful connection to the database")
	}

	return DB, err
}

func (r *Repository) GetContainers() ([]models.Container, error) {
	var containers []models.Container
	err := r.DB.Select(&containers,
		"SELECT id, ip, status, last_ping_time, ping_duration FROM containers")
	log.Println("containers from db = ", containers)
	if err != nil {
		log.Println("err while getting containers from db = ", err)
		return nil, err
	}
	return containers, nil
}

func (r *Repository) AddContainer(containerJSON []byte) error {
	var container models.Container
	err := json.Unmarshal(containerJSON, &container)
	if err != nil {
		log.Fatalln("error with unmarshalling containerData: ", err)
		return err
	}
	log.Println("unmarshalled container = ", container)

	existingID := -1
	err = r.DB.QueryRow("SELECT id FROM containers WHERE ip = $1", container.IP).Scan(&existingID)
	log.Println("existingID: ", existingID)
	log.Println("err:", err)

	// если такой контейнер уже есть
	if err == nil {
		_, err = r.DB.Exec("UPDATE containers SET status = $1, last_ping_time = $2, ping_duration = $3 WHERE ip = $4",
			container.Status, container.LastPingTime, container.PingDuration, container.IP)
		if err != nil {
			log.Fatalln("error with UPDATE containers in the DB: ", err)
			return err
		}
		return nil
	}

	// если такого контейнера нет
	if errors.Is(err, sql.ErrNoRows) {
		_, err = r.DB.Exec("INSERT INTO containers (ip, status, last_ping_time, ping_duration) VALUES ($1, $2, $3, $4)",
			container.IP, container.Status, container.LastPingTime, container.PingDuration)
		if err != nil {
			log.Fatalln("error with INSERT container in the DB: ", err)
			return err
		}
		return nil
	}

	// Если произошла другая ошибка
	log.Fatalln("another type of error: ", err)
	return err
}
