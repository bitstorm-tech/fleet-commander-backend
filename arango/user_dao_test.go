package arango

import (
	"testing"
)

func TestInsertUser(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	if err := SetupFleetCommanderDatabase(true); err != nil {
		t.Errorf("Error while setup database")
	}

}

func TestInsertInvalidUser(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	if err := SetupFleetCommanderDatabase(true); err != nil {
		t.Errorf("Error while setup database")
	}

	if err := InsertNewUser(nil); err == nil {
		t.Error("Expected error but got nil")
	}

	user := new(User)
	if err := InsertNewUser(user); err == nil {
		t.Errorf("Expected error but got nil")
	}
}
