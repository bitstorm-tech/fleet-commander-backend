package game

import "time"

type Resources struct {
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}
