package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddExoPlanet(t *testing.T) {

	requestBody := []byte(`{"name":"Test Exoplanet", "description":"Test Description", "distance":500, "radius":2.5, "type":"GasGiant"}`)
	req, err := http.NewRequest("POST", "/exoplanets", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(AddExoPlanet)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestListExoPlanets(t *testing.T) {

	req, err := http.NewRequest("GET", "/exoplanets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(ListExoPlanets)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetExoPlanetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/exoplanets/non-existing-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetExoPlanetByID)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestUpdateExoPlanet(t *testing.T) {
	requestBody := []byte(`{"name":"Updated Exoplanet", "description":"Updated Description", "distance":500, "radius":2.5, "type":"GasGiant"}`)
	req, err := http.NewRequest("PUT", "/exoplanets/1", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(UpdateExoPlanet)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestDeleteExoPlanet(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/exoplanets/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeleteExoPlanet)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

}

func TestFuelEstimation(t *testing.T) {

	req, err := http.NewRequest("GET", "/exoplanets/1/fuel?crew=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FuelEstimation)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
