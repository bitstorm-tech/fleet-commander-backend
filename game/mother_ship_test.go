package game

import "testing"

var testMotherShip = MotherShip{
	TitaniumPerMinute: 10,
	FuelPerMinute:     10,
	EnergyPerMinute:   10,
}

var motherShipCalcTestData = []struct {
	deltaMinutes     float64
	expectedTitanium int
	expectedFuel     int
	expectedEnergy   int
}{
	{0.001, 0, 0, 0},
	{0.01, 0, 0, 0},
	{0.1, 1, 1, 1},
	{0.15, 1, 1, 1},
	{0.25, 2, 2, 2},
	{0.29, 2, 2, 2},
	{0.5, 5, 5, 5},
	{1, 10, 10, 10},
}

func TestCalcResources2(t *testing.T) {
	for _, data := range motherShipCalcTestData {
		result := testMotherShip.CalcResources(data.deltaMinutes)

		if result.Titanium != data.expectedTitanium {
			t.Errorf("(deltaMinutes = %v) expected %v, got %v", data.deltaMinutes, data.expectedTitanium, result.Titanium)
		}

		if result.Fuel != data.expectedFuel {
			t.Errorf("(deltaMinutes = %v) expected %v, got %v", data.deltaMinutes, data.expectedFuel, result.Fuel)
		}

		if result.Energy != data.expectedEnergy {
			t.Errorf("(deltaMinutes = %v) expected %v, got %v", data.deltaMinutes, data.expectedEnergy, result.Energy)
		}
	}
}
