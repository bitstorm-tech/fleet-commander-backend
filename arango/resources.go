package arango

import (
	"time"
)

const (
	resourcesCollectionName = "resources"
)

type Resources struct {
	Key        string    `json:"_key,omitempty"`
	ID         string    `json:"_id,omitempty"`
	Titanium   int       `json:"titanium"`
	Fuel       int       `json:"fuel"`
	Energy     int       `json:"energy"`
	LastUpdate time.Time `json:"lastUpdate"`
}

func (Resources) collection() string {
	return resourcesCollectionName
}

func (r *Resources) key() string {
	return r.Key
}

func (r *Resources) id() string {
	return r.ID
}

func (r *Resources) setKey(key string) {
	r.Key = key
	r.ID = r.collection() + "/" + key
}
