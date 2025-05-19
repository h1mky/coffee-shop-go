package http

import (
	"coffee-shop/db"
	"net/http"
)

func GetSingleCardHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	params, err := getParams(r)
	if err != nil {
		return
	}
	card, err := db.GetSingleCard(ctx, params)
	if err != nil {
		http.Error(w, "Failed to Fetch data ", http.StatusInternalServerError)
		return
	}
	JsonEncode(w, card)

}
