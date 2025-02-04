package main

import (
	"encoding/json"
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
	_, err = db.Exec("INSERT INTO containers (ip, status, last_ping_time) VALUES ($1, $2, $3)",
		container.IP, container.Status, container.LastPingTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
