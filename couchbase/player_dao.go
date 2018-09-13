package couchbase

import (
	"github.com/bugjoe/fleet-commander-backend/game"
	"github.com/pkg/errors"
	"gopkg.in/couchbase/gocb.v1"
	"log"
	"strings"
)

func InsertNewPlayerWithLogin(l game.Login) error {
	passwordHash, err := l.PasswordHash()
	if err != nil {
		return errors.WithStack(err)
	}

	l.Password = passwordHash
	log.Println("insert new player with login:", l)

	bucket, err := getBucket("fc-player")
	if err != nil {
		return errors.WithStack(err)
	}

	p := game.NewPlayer()
	p.Login = l
	_, err = bucket.Insert(strings.ToLower(l.Email), p, 0)
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
