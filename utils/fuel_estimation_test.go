package utils

import (
	"testing"

	"github.com/Kratos-28/ExoPlanet/models"
)

func TestCalculateFuel(t *testing.T) {

	exoplanet := models.Exoplanet{
		Distance: 500,
		Radius:   1.11,
		Mass:     float64Ptr(1.0),
		Type:     "GasGiant",
	}

	fuelCost, err := CalculateFuel(exoplanet, 5)
	if err != nil {
		t.Errorf("calculateFuel returned an error: %v", err)
	}

	expectedFuelCost := 607.2281640000002

	if fuelCost != expectedFuelCost {
		t.Errorf("calculateFuel returned wrong fuel cost: got %v want %v", fuelCost, expectedFuelCost)
	}
}
func float64Ptr(f float64) *float64 {
	return &f
}
