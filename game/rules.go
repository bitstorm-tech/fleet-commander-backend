package game

var ActiveRules Rules

type Rules struct {
	TitaniumHarvester ShipRule `json:"titaniumHarvester"`
	FuelHarvester     ShipRule `json:"fuelHarvester"`
}

type ShipRule struct {
	TitaniumCost     int `json:"titaniumCost"`
	FuelCost         int `json:"fuelCost"`
	HarvestPerMinute int `json:"harvestPerMinute"`
	HitPoints        int `json:"hitPoints"`
}

func (r *ShipRule) WithTitaniumCost(c int) *ShipRule {
	r.TitaniumCost = c
	return r
}

func (r *ShipRule) WithFuelCost(c int) *ShipRule {
	r.FuelCost = c
	return r
}

func (r *ShipRule) WithHarvestPerMinute(hpm int) *ShipRule {
	r.HarvestPerMinute = hpm
	return r
}

func (r *ShipRule) WithHitPoints(hp int) *ShipRule {
	r.HitPoints = hp
	return r
}

func NewShipRule() *ShipRule {
	return new(ShipRule)
}

func (Rules) BucketName() string {
	return "fc-internal"
}

func (Rules) ID() string {
	return "rules"
}
