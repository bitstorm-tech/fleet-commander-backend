package couchbase

import (
	"github.com/bugjoe/fleet-commander-backend/game"
	"github.com/pkg/errors"
)

func GetGameRules() (game.Rules, error) {
	rules := new(game.Rules)
	err := Get(rules)
	if err != nil {
		return game.Rules{}, errors.WithStack(err)
	}

	return *rules, nil
}
