package main

import (
	"fmt"
	"github.com/bugjoe/fleet-commander-backend/couchbase"
	"github.com/bugjoe/fleet-commander-backend/game"
)

var rules = game.Rules{
	TitaniumHarvester: *game.NewShipRule().WithHitPoints(500).WithHarvestPerMinute(05).WithTitaniumCost(100).WithFuelCost(050),
	FuelHarvester:     *game.NewShipRule().WithHitPoints(500).WithHarvestPerMinute(05).WithTitaniumCost(100).WithFuelCost(050),
}

func main() {
	fmt.Println("Remove old game rules")
	couchbase.Delete(rules)

	fmt.Println("Inserting game rules")
	couchbase.Insert(rules)
}
