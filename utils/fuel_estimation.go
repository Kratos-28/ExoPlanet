package utils

import (
	"errors"
	"math"

	"github.com/Kratos-28/ExoPlanet/models"
)

func CalculateFuel(exoplanet models.Exoplanet, crewCapacity int) (float64, error) {
	var g float64
	switch exoplanet.Type {
	case "GasGiant":
		g = 0.5 / math.Pow(exoplanet.Radius, 2)

	case "Terrestial":
		if exoplanet.Mass == nil {
			return 0, errors.New("mass is required for terrestial exoplanets")

		}
		g = *exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	default:
		return 0, errors.New("unknown exoplanet type")
	}

	fuelCost := exoplanet.Distance / (math.Pow(g, 2) * float64(crewCapacity))
	return fuelCost, nil
}
