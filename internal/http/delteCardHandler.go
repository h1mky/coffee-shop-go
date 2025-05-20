package http

import (
	"coffee-shop/db"
	"net/http"
)

func DeleteCardHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params, err := getParams(r)
	if err != nil {
		http.Error(w, "failed to find params", http.StatusBadRequest)
		return
	}

	if err := db.DeleteCard(ctx, params); err != nil {
		http.Error(w, "Failed to delete data ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
