package arango

// Persistable is the interface that is shared by all classes that can be
// persisted in the ArangoDB database.
type Persistable interface {
	collection() string
	key() string
	id() string
	setKey(key string)
}
