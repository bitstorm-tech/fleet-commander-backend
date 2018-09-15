package game

import "time"

type Resources struct {
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}

// Add resources to this resources and returns a new Resources object.
// So, this function does not mutate the receiving resources.
func (r Resources) Add(other Resources) Resources {
	return Resources{
		Titanium: r.Titanium + other.Titanium,
		Fuel:     r.Fuel + other.Fuel,
		Energy:   r.Energy + other.Energy,
	}
}
