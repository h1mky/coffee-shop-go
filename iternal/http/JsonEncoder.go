package http

import (
	"encoding/json"
	"net/http"
)

func JsonEncode(w http.ResponseWriter, cards any) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(cards)
	if err != nil {
		http.Error(w, "Failed to Encode Json", http.StatusInternalServerError)
	}
}
