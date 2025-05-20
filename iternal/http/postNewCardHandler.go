package http

import (
	"coffee-shop/db"
	"coffee-shop/models"
	validate "coffee-shop/validator"

	"net/http"
)

func PostNewCardHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var card models.CoffeeCardInput

	if err := jsonParse(w, r, &card); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := validate.ValidateStruct(card); err != nil {
		http.Error(w, "error on validate json", http.StatusBadRequest)
		return
	}

	if err := db.PostNewCard(ctx, card); err != nil {
		http.Error(w, "failed to save data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
