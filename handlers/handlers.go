package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kratos-28/ExoPlanet/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
		return
	}
	exoplanet.ID = uuid.NewString()
	store.Lock()
	store.Exoplanets[exoplanet.ID] = exoplanet
	store.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoPlanets(w http.ResponseWriter, r *http.Request) {
	store.RLock()
	defer store.RUnlock()
	exoplanets := make([]models.Exoplanet, 0, len(store.Exoplanets))
	for _, planet := range store.Exoplanets {
		exoplanets = append(exoplanets, planet)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)

}

func GetExoPlanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	store.RLock()
	defer store.RUnlock()
	exoplanet, ok := store.Exoplanets[id]
	if !ok {
		respondWithError(w, http.StatusNotFound, models.ErrNotFound.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoPlanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var exoplanet models.Exoplanet

	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	if err := validateExoplanet(&exoplanet); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid exoplanet data")
		return
	}
	store.Lock()
	defer store.Unlock()
	_, ok := store.Exoplanets[id]
	if !ok {
		respondWithError(w, http.StatusNotFound, models.ErrNotFound.Error())
		return
	}
	exoplanet.ID = id
	store.Exoplanets[id] = exoplanet
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}

func DeleteExoPlanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	store.Lock()
	defer store.Unlock()
	if _, ok := store.Exoplanets[id]; !ok {
		respondWithError(w, http.StatusNotFound, models.ErrNotFound.Error())
		return
	}
	delete(store.Exoplanets, id)
	w.WriteHeader(http.StatusNoContent)
}
