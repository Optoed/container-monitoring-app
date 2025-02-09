package containerHandler

import (
	"backend/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler) GetContainers(w http.ResponseWriter, r *http.Request) {
	containers, err := h.Service.GetContainers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containers)
}
