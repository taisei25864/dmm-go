package timelines

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// )

// func (h *handler) GetPublicTimeline(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	tmp_maxID := r.URL.Query().Get("max_id")
// 	tmp_sinceID := r.URL.Query().Get("since_id")
// 	tmp_limit := r.URL.Query().Get("limit")

// 	maxID, _ := strconv.ParseInt(tmp_maxID, 10, 64)
// 	sinceID, _ := strconv.ParseInt(tmp_sinceID, 10, 64)
// 	limit, _ := strconv.ParseInt(tmp_limit, 10, 64)

// 	status, err := h.tr.GetPublicTimeline(ctx, maxID, sinceID, limit)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(status); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }
