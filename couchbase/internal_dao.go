package couchbase

import (
	"github.com/pkg/errors"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/game"
)

func GetGameRules() (game.Rules, error) {
	rules := new(game.Rules)
	err := Get(rules)
	if err != nil {
		return game.Rules{}, errors.WithStack(err)
	}

	return *rules, nil
}
