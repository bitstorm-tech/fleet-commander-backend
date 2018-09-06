package arango

import (
	"github.com/pkg/errors"
)

func CreateDocument(p Persistable) error {
	collection, err := getCollection(p.collection())
	if err != nil {
		return errors.WithStack(err)
	}

	meta, err := collection.CreateDocument(nil, p)
	if err != nil {
		return errors.WithStack(err)
	}

	p.setKey(meta.Key)

	return nil
}

func RemoveDocument(p Persistable) error {
	collection, err := getCollection(p.collection())
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = collection.RemoveDocument(nil, p.key())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
