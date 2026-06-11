package handlers

import (
	"encoding/json"
	"net/http"
	"shop/internal/api"
)

func (h *Handler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user api.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Couldn't decode json body", http.StatusBadRequest)
		return
	}

	if err := h.repository.CreateUsesr(user); err != nil {
		http.Error(w, "Internal Error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
