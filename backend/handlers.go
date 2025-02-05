package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

func GetContainers(w http.ResponseWriter, r *http.Request) {
	var containers []Container
	err := db.Select(&containers, "SELECT * FROM containers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}

func AddContainer(w http.ResponseWriter, r *http.Request) {
	var container Container
	err := json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var existingID int
	err = db.QueryRow("SELECT id FROM containers WHERE ip = $1", container.IP).Scan(&existingID)

	// если такой контейнер уже есть
	if err == nil {
		_, err = db.Exec("UPDATE containers SET status = $1, last_ping_time = $2 WHERE ip = $3",
			container.Status, container.LastPingTime, container.IP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	// если такого контейнера нет
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.Exec("INSERT INTO containers (ip, status, last_ping_time) VALUES ($1, $2, $3)",
			container.IP, container.Status, container.LastPingTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

	// Если произошла другая ошибка
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
