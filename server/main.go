package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kratos-28/ExoPlanet/handlers"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Printf("Started application at port :8080 \n")
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", handlers.AddExoPlanet).Methods("POST")
	r.HandleFunc("/exoplanets", handlers.ListExoPlanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.GetExoPlanetByID).Methods("GET")
	// r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoPlanet).Methods("PUT")
	// r.HandleFunc("/exoplanets/{id}", handlers.DeleteExoPlanet).Methods("DELETE")
	// r.HandleFunc("/exoplanets/{id}", handlers.FuelEstimation).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
