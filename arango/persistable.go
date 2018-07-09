package arango

type Persistable interface {
	Collection() string
	Key() string
	ID() string
}
