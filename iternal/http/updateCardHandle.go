package http

import (
	"coffee-shop/db"
	"coffee-shop/models"
	"net/http"
)

func UpdateCardHandle(w http.ResponseWriter, r *http.Request) {

	var card models.CoffeeCardInput

	ctx := r.Context()

	params, err := getParams(r)
	if err != nil {
		http.Error(w, "failed to find params", http.StatusBadRequest)
		return
	}
	if err := jsonParse(w, r, &card); err != nil {
		http.Error(w, "ERROR with Json parse", http.StatusBadRequest)
		return
	}

	if err := db.UpdateCard(ctx, params, card); err != nil {
		http.Error(w, "Failed to update card", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
