package main

import (
	"fmt"
	"github.com/bugjoe/fleet-commander-backend/couchbase"
	"github.com/bugjoe/fleet-commander-backend/game"
)

var rules = game.Rules{
	TitaniumHarvester: []game.ShipRule{
		*game.NewShipRule().WithLevel(0).WithHitPoints(500).WithHarvestPerMinute(05).WithTitaniumCost(100).WithFuelCost(050),
		*game.NewShipRule().WithLevel(1).WithHitPoints(550).WithHarvestPerMinute(10).WithTitaniumCost(150).WithFuelCost(100),
		*game.NewShipRule().WithLevel(2).WithHitPoints(600).WithHarvestPerMinute(15).WithTitaniumCost(200).WithFuelCost(150),
		*game.NewShipRule().WithLevel(3).WithHitPoints(650).WithHarvestPerMinute(20).WithTitaniumCost(250).WithFuelCost(200),
		*game.NewShipRule().WithLevel(4).WithHitPoints(700).WithHarvestPerMinute(25).WithTitaniumCost(300).WithFuelCost(250),
	},
	FuelHarvester: []game.ShipRule{
		*game.NewShipRule().WithLevel(0).WithHitPoints(500).WithHarvestPerMinute(05).WithTitaniumCost(100).WithFuelCost(050),
		*game.NewShipRule().WithLevel(1).WithHitPoints(550).WithHarvestPerMinute(10).WithTitaniumCost(150).WithFuelCost(100),
		*game.NewShipRule().WithLevel(2).WithHitPoints(600).WithHarvestPerMinute(15).WithTitaniumCost(200).WithFuelCost(150),
		*game.NewShipRule().WithLevel(3).WithHitPoints(650).WithHarvestPerMinute(20).WithTitaniumCost(250).WithFuelCost(200),
		*game.NewShipRule().WithLevel(4).WithHitPoints(700).WithHarvestPerMinute(25).WithTitaniumCost(300).WithFuelCost(250),
	},
}

func main() {
	fmt.Println("Remove old game rules")
	couchbase.Delete(rules)

	fmt.Println("Inserting game rules")
	couchbase.Insert(rules)
}
