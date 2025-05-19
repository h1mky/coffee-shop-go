package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func getParams(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.Atoi(idStr)
}
