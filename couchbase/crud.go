package couchbase

import (
	"github.com/pkg/errors"
	"gopkg.in/couchbase/gocb.v1"
)

const (
	dbConnection = "couchbase://localhost"
	dbUser       = "Administrator"
	dbPassword   = "password"
)

func getCluster() (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(dbConnection)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = cluster.Authenticate(gocb.PasswordAuthenticator{Username: dbUser, Password: dbPassword})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cluster, nil
}

func getBucket(name string) (*gocb.Bucket, error) {
	cluster, err := getCluster()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bucket, err := cluster.OpenBucket(name, "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bucket, nil
}

func Insert(p Persistable) error {
	bucket, err := getBucket(p.BucketName())
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = bucket.Insert(p.ID(), &p, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Delete(p Persistable) error {
	bucket, err := getBucket(p.BucketName())
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = bucket.Remove(p.ID(), 0)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func Get(p Persistable) error {
	bucket, err := getBucket(p.BucketName())
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = bucket.Get(p.ID(), p)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

type Persistable interface {
	BucketName() string
	ID() string
}
