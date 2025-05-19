package http

import (
	"coffee-shop/db"

	"net/http"
)

func GetAllCardHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	cards, err := db.GetAllCard(ctx)
	if err != nil {
		http.Error(w, "Failed to Fetch data ", http.StatusInternalServerError)
		return
	}

	JsonEncode(w, cards)
}
