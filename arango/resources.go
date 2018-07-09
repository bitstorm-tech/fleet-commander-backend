package arango

import (
	"time"
)

type Resources struct {
	Key_       string    `json:"_key"`
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}

func (Resources) Collection() string {
	return CollectionResource
}

func (resources *Resources) Key() string {
	return resources.Key_
}

func (resources *Resources) ID() string {
	if len(resources.Key()) == 0 {
		return ""
	}

	return resources.Collection() + "/" + resources.Key()
}
