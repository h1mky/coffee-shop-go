package http

import (
	"coffee-shop/models"
	validate "coffee-shop/validator"
	"net/http"
)

func PostNewCard(w http.ResponseWriter, r *http.Request) {

	var card models.CoffeeCardInput

	err := jsonParse(w, r, &card)
	if err != nil {
		http.Error(w, "failed parse json", http.StatusInternalServerError)
		return
	}
	err2 := validate.ValidateStruct(card)
	if err2 != nil {
		http.Error(w, "error on validate json", http.StatusBadRequest)
		return
	}

}
