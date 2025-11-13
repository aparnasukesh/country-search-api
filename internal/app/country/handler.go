package country

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type CountryHandler struct {
	service *CountryService
}

func NewCountryHandler(service *CountryService) *CountryHandler {
	return &CountryHandler{service: service}
}

func (h *CountryHandler) SearchCountry(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing 'name' parameter", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	country, err := h.service.GetCountry(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}
