package models

import (
	"errors"
	"sync"
)

type Exoplanet struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Distance    float64  `json:"distance"`
	Radius      float64  `json:"radius"`
	Type        string   `json:"type"`
	Mass        *float64 `json:"mass,omitempty"`
}

type ExoplanetStore struct {
	sync.RWMutex
	Exoplanets map[string]Exoplanet
}

func NewExoPlanetStore() *ExoplanetStore {
	return &ExoplanetStore{
		Exoplanets: make(map[string]Exoplanet),
	}
}

var (
	ErrNotFound = errors.New("exoplanet not found")
	ErrInvalid  = errors.New("invalid Exoplanet data")
)
