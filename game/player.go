package game

import (
	"time"
)

// Player is the structure represents a player from the database
type Player struct {
	Login      Login      `json:"login"`
	Resources  Resources  `json:"resources"`
	Ships      Ships      `json:"ships"`
	MotherShip MotherShip `json:"motherShip"`
}

// ActualResources calculates the actual resources of the player and returns them.
// These resources are the resources from the mother ship plus the resources from the
// harvesters since the last update time.
// The actual resources are only temporary and will not be stored in the database.
func (p Player) ActualResources() Resources {
	deltaTimeMinutes := int(time.Now().Sub(p.Resources.LastUpdate).Seconds() / 60)

	return p.calcResourcesFromMotherShip(deltaTimeMinutes).Add(p.calcResourcesFromHarvester(deltaTimeMinutes))
}

func (p Player) calcResourcesFromMotherShip(d int) Resources {
	return Resources{
		Titanium: d * p.MotherShip.TitaniumPerMinute,
		Fuel:     d * p.MotherShip.FuelPerMinute,
		Energy:   d * p.MotherShip.EnergyPerMinute,
	}
}

func (p Player) calcResourcesFromHarvester(d int) Resources {
	return Resources{
		Titanium: d * p.Ships.TitaniumHarvester,
		Fuel:     d * p.Ships.FuelHarvester,
	}
}

// NewPlayer returns a new player which resource last update time is
// set to time.Now()
func NewPlayer() Player {
	p := Player{}
	p.Resources.LastUpdate = time.Now()
	p.MotherShip.EnergyPerMinute = 1
	p.MotherShip.TitaniumPerMinute = 10
	p.MotherShip.FuelPerMinute = 10
	p.Ships.TitaniumHarvester = 10
	p.Ships.FuelHarvester = 10

	return p
}

func (Player) BucketName() string {
	return "fc-player"
}

func (p Player) ID() string {
	return p.Login.Email
}
