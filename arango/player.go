package arango

import (
	"crypto"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

const (
	playerCollectionName = "player"
)

// Player is the structure represents a player from the database
type Player struct {
	Key       string    `json:"_key,omitempty"`
	ID        string    `json:"_id,omitempty"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Resources Resources `json:"resources"`
	Ships     Ships     `json:"ships"`
}

// Resources of a player
type Resources struct {
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Ships struct {
	TitaniumHarvesterAmount int `json:"titaniumHarvesterAmount"`
	FuelHarvesterAmount     int `json:"fuelHarvesterAmount"`
	EnergyHarvesterAmount   int `json:"energyHarvesterAmount"`
}

// PasswordHash returns the players password as hex encoded SHA-512 hash string
func (p *Player) PasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(p.Password)); err != nil {
		return "", errors.WithStack(err)
	}

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
}

func (Player) collection() string {
	return playerCollectionName
}

func (p *Player) key() string {
	return p.Key
}

func (p *Player) id() string {
	return p.ID
}

func (p *Player) setKey(key string) {
	p.Key = key
	p.ID = p.collection() + "/" + key
}
