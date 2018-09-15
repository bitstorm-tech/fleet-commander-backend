package game

import (
	"math"
)

type Ships struct {
	TitaniumHarvester int `json:"titaniumHarvester"`
	FuelHarvester     int `json:"fuelHarvester"`
}

func (s Ships) CalcResources(deltaMinutes float64) Resources {
	titaniumPerMinute := ActiveRules.TitaniumHarvester.HarvestPerMinute
	fuelPerMinute := ActiveRules.FuelHarvester.HarvestPerMinute

	return Resources{
		Titanium: int(math.Round(deltaMinutes * float64(s.TitaniumHarvester*titaniumPerMinute))),
		Fuel:     int(math.Round(deltaMinutes * float64(s.FuelHarvester*fuelPerMinute))),
	}
}
