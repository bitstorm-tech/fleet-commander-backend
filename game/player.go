package game

import (
	"crypto"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

// Player is the structure represents a player from the database
type Player struct {
	Login      Login      `json:"login"`
	Resources  Resources  `json:"resources"`
	Ships      Ships      `json:"ships"`
	MotherShip MotherShip `json:"motherShip"`
}

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Resources struct {
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Ships struct {
	TitaniumHarvester int `json:"titaniumHarvester"`
	FuelHarvester     int `json:"fuelHarvester"`
}

type MotherShip struct {
	EnergyPerMinute   int `json:"energyPerMinute"`
	TitaniumPerMinute int `json:"titaniumPerMinute"`
	FuelPerMinute     int `json:"fuelPerMinute"`
}

// PasswordHash returns the players password as hex encoded SHA-512 hash string
func (l Login) PasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(l.Password)); err != nil {
		return "", errors.WithStack(err)
	}

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
}

// ActualResources calculates the actual resources of the player and returns them.
// The actual resources are only temporary and will not be stored in the database.
func (p Player) ActualResources() Resources {
	deltaTimeSeconds := int(time.Now().Sub(p.Resources.LastUpdate).Seconds())
	p.Resources.Titanium += deltaTimeSeconds
	p.Resources.Fuel += deltaTimeSeconds
	p.Resources.Energy += deltaTimeSeconds
	return p.Resources
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
