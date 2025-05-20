package http

import (
	"coffee-shop/models"
	"net/http"
)

func PostNewCard(w http.ResponseWriter, r *http.Request) {

	var card models.CoffeeCard

	err := jsonParse(w, r, &card)
	if err != nil {
		http.Error(w, "failed parse json", http.StatusInternalServerError)
		return
	}

}
