package game

type MotherShip struct {
	EnergyPerMinute   int `json:"energyPerMinute"`
	TitaniumPerMinute int `json:"titaniumPerMinute"`
	FuelPerMinute     int `json:"fuelPerMinute"`
}

func (m MotherShip) CalcResources(deltaMinutes float64) Resources {
	return Resources{
		Titanium: int(deltaMinutes * float64(m.TitaniumPerMinute)),
		Fuel:     int(deltaMinutes * float64(m.FuelPerMinute)),
		Energy:   int(deltaMinutes * float64(m.EnergyPerMinute)),
	}
}
