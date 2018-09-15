package game

import "testing"

func TestAdd(t *testing.T) {
	a := Resources{
		Titanium: 1,
		Fuel:     1,
		Energy:   1,
	}

	b := Resources{
		Titanium: 2,
		Fuel:     2,
		Energy:   2,
	}

	result := a.Add(b)

	if result.Titanium != 3 {
		t.Error("expected 3, got ", result.Titanium)
	}

	if result.Fuel != 3 {
		t.Error("expected 3, got ", result.Fuel)
	}

	if result.Energy != 3 {
		t.Error("expected 3, got ", result.Energy)
	}
}
