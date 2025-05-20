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

	err := jsonParse(w, r, &card)
	if err != nil {
		http.Error(w, "failed parse json", http.StatusInternalServerError)
		return
	}
	err = validate.ValidateStruct(card)
	if err != nil {
		http.Error(w, "error on validate json", http.StatusBadRequest)
		return
	}

	err = db.PostNewCard(ctx, card)
	if err != nil {
		http.Error(w, "failed to save data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
