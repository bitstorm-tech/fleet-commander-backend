package game

import (
	"crypto"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

// Player is the structure represents a player from the database
type Player struct {
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Resources Resources `json:"resources"`
	Ships     Ships     `json:"ships"`
}

type Ships struct {
	TitaniumHarvester int `json:"titaniumHarvester"`
	FuelHarvester     int `json:"fuelHarvester"`
	EnergyHarvester   int `json:"energyHarvester"`
}

// PasswordHash returns the players password as hex encoded SHA-512 hash string
func (p Player) PasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(p.Password)); err != nil {
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
	return p
}

func (Player) BucketName() string {
	return "fc-player"
}

func (p Player) ID() string {
	return p.Email
}
