package arango

import (
	"testing"
)

func TestInsertUser(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	SetupFleetCommanderDatabase(true)
}

func TestInsertInvalidUser(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	SetupFleetCommanderDatabase(true)

	err := InsertNewUser(nil)
	if err == nil {
		t.Error("Expected error but got nil")
	}

	user := new(User)
	err = InsertNewUser(user)
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}
