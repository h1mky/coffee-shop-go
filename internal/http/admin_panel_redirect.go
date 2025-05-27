package http

import (
	"net/http"
	"os"
)

type UserStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminPanelRedirect(w http.ResponseWriter, r *http.Request) {
	valideName := os.Getenv("USERNAME")
	validePass := os.Getenv("ADMIN_PASSWORD")

	var req UserStruct
	if err := jsonParse(w, r, &req); err != nil {
		http.Error(w, "Error on parsing json", http.StatusBadRequest)
		return
	}

	if valideName == req.Username && validePass == req.Password {
		w.WriteHeader(http.StatusNoContent)
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
