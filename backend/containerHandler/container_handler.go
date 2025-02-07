package containerHandler

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Handler struct {
	DB *sqlx.DB
}

func (h *Handler) GetContainers(w http.ResponseWriter, r *http.Request) {
	var containers []models.Container
	err := h.DB.Select(&containers,
		"SELECT id, ip, status, last_ping_time, EXTRACT(EPOCH FROM ping_duration)::BIGINT AS ping_duration FROM containers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}

func (h *Handler) AddContainer(w http.ResponseWriter, r *http.Request) {
	var container models.Container
	err := json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//if h.DB == nil {
	//	log.Println("Fail connection to db")
	//}

	var existingID int
	err = h.DB.QueryRow("SELECT id FROM containers WHERE ip = $1", container.IP).Scan(&existingID)

	//log.Println("existingID: ", existingID)
	//log.Println("err:", err)

	// если такой контейнер уже есть
	if err == nil {
		_, err = h.DB.Exec("UPDATE containers SET status = $1, last_ping_time = $2, ping_duration = $3 WHERE ip = $4",
			container.Status, container.LastPingTime, container.PingDuration, container.IP)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	// если такого контейнера нет
	if errors.Is(err, sql.ErrNoRows) {
		_, err = h.DB.Exec("INSERT INTO containers (ip, status, last_ping_time, ping_duration) VALUES ($1, $2, $3, $4)",
			container.IP, container.Status, container.LastPingTime, container.PingDuration)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}

	// Если произошла другая ошибка
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
