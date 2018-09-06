package arango

import (
	"fmt"

	"github.com/pkg/errors"
)

// SetupFleetCommanderDatabase creates a new fleet-commander database with all its
// collections. If drop is true, a already existing fleet-commander database will be
// dropped.
func SetupFleetCommanderDatabase(drop bool) error {
	fmt.Println("Getting database ...")
	database, err := getDatabase()
	if err != nil {
		return errors.WithStack(err)
	}

	if drop {
		fmt.Println("Dropping database ...")
		err = database.Remove(nil)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	fmt.Println("Getting client ...")
	client, err := getClient()
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println("Creating database ...")
	_, err = client.CreateDatabase(nil, databaseName, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println("Creating player collection ...")
	_, err = database.CreateCollection(nil, playerCollectionName, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
