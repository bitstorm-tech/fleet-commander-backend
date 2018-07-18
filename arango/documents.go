package arango

import (
	"github.com/pkg/errors"
)

func CreateDocument(p Persistable) error {
	collection, err := getCollection(p.collection())
	if err != nil {
		return errors.Wrapf(err, "error while getting collection : '%s'", p.collection())
	}

	meta, err := collection.CreateDocument(nil, p)
	if err != nil {
		return errors.Wrapf(err, "error while creating document: %+v", p)
	}

	p.setKey(meta.Key)

	return nil
}

func RemoveDocument(p Persistable) error {
	collection, err := getCollection(p.collection())
	if err != nil {
		return errors.Wrapf(err, "error while removing document with key: %s", p.key())
	}

	_, err = collection.RemoveDocument(nil, p.key())
	if err != nil {
		return errors.Wrapf(err, "error while removing document: %s", p.key())
	}

	return nil
}
