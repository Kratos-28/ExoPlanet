package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kratos-28/ExoPlanet/models"
	"github.com/google/uuid"
)

var store = models.NewExoPlanetStore()

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func validateExoplanet(exoplanet *models.Exoplanet) error {
	if exoplanet.Name == "" || exoplanet.Description == "" {
		return models.ErrInvalid
	}
	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		return models.ErrInvalid
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return models.ErrInvalid
	}
	if exoplanet.Type == "Terrestrial" {
		if exoplanet.Mass == nil || *exoplanet.Mass <= 0.1 || *exoplanet.Mass >= 10 {
			return models.ErrInvalid
		}
	}
	if exoplanet.Type != "GasGiant" && exoplanet.Type != "Terrestrial" {
		return models.ErrInvalid
	}
	return nil
}
func AddExoPlanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validateExoplanet(&exoplanet); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid exoplanet data")
	}
	exoplanet.ID = uuid.NewString()
	store.Lock()
	store.Exoplanets[exoplanet.ID] = exoplanet
	store.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}
