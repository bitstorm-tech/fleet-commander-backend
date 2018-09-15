package game

import "testing"

var testShips = Ships{
	TitaniumHarvester: 10,
	FuelHarvester:     10,
}

var shipCalcTestData = []struct {
	deltaMinutes     float64
	expectedTitanium int
	expectedFuel     int
}{
	{0.001, 0, 0},
	{0.01, 1, 1},
	{0.1, 10, 10},
	{0.15, 15, 15},
	{0.25, 25, 25},
	{0.29, 29, 29},
	{0.5, 50, 50},
	{1, 100, 100},
}

func init() {
	ActiveRules = Rules{
		TitaniumHarvester: *NewShipRule().WithHarvestPerMinute(10),
		FuelHarvester:     *NewShipRule().WithHarvestPerMinute(10),
	}
}

func TestCalcResources(t *testing.T) {
	for _, testData := range shipCalcTestData {
		result := testShips.CalcResources(testData.deltaMinutes)

		if result.Titanium != testData.expectedTitanium {
			t.Errorf("(deltaMinutes = %v) expected %v, got %v", testData.deltaMinutes, testData.expectedTitanium, result.Titanium)
		}

		if result.Fuel != testData.expectedFuel {
			t.Errorf("(deltaMinutes = %v) expected %v, got %v", testData.deltaMinutes, testData.expectedFuel, result.Fuel)
		}

		if result.Energy != 0 {
			t.Error("expected energy to be 0, got ", result.Energy)
		}
	}
}
