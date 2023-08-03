package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) FindStatusByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	temp_id := chi.URLParam(r, "id")

	id, _ := strconv.ParseInt(temp_id, 10, 64)

	status, err := h.sr.FindStatusByID(ctx, id)
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
