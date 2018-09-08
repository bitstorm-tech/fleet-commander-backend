package couchbase

import (
	"github.com/pkg/errors"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/game"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"strings"
)

func InsertNewPlayer(p game.Player) error {
	passwordHash, err := p.PasswordHash()
	if err != nil {
		return errors.WithStack(err)
	}

	p.Password = passwordHash
	log.Println("insert new player:", p)

	bucket, err := getBucket("fc-player")
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = bucket.Insert(strings.ToLower(p.Email), p, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func GetPlayerByEmail(email string) (*game.Player, error) {
	bucket, err := getBucket("fc-player")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	p := new(game.Player)
	_, err = bucket.Get(strings.ToLower(email), p)
	if err != nil {
		if gocb.IsKeyNotFoundError(err) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}

	return p, nil
}
