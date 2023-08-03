package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type AddRequest struct {
	Status string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.Content = req.Status
	account := auth.AccountOf(r)
	status.AccountID = account.ID

	if err := h.sr.Create(r.Context(), status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, err := h.sr.FindStatusByID(r.Context(), status.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
