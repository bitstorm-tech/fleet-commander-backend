package arango

import (
	"testing"
)

func TestInsertPlayer(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	if err := SetupFleetCommanderDatabase(true); err != nil {
		t.Errorf("Error while setup database")
	}

}

func TestInsertInvalidPlayer(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	if err := SetupFleetCommanderDatabase(true); err != nil {
		t.Errorf("Error while setup database")
	}

	if err := InsertNewPlayer(nil); err == nil {
		t.Error("Expected error but got nil")
	}

	p := new(Player)
	if err := InsertNewPlayer(p); err == nil {
		t.Errorf("Expected error but got nil")
	}
}
