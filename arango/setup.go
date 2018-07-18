package arango

import (
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/pkg/errors"
)

// SetupFleetCommanderDatabase creates a new fleet-commander database with all its
// collections. If drop is true, a already existing fleet-commander database will be
// dropped.
func SetupFleetCommanderDatabase(drop bool) error {
	fmt.Println("Getting database ...")
	database, err := getDatabase()
	if err != nil {
		return errors.Wrap(err, "can't get database")
	}

	if drop {
		fmt.Println("Dropping database ...")
		err = database.Remove(nil)
		if err != nil {
			return errors.Wrap(err, "can't drop database")
		}
	}

	fmt.Println("Getting client ...")
	client, err := getClient()
	if err != nil {
		return errors.Wrap(err, "can't get client")
	}

	fmt.Println("Creating database ...")
	_, err = client.CreateDatabase(nil, databaseName, nil)
	if err != nil {
		return errors.Wrap(err, "can't create database")
	}

	fmt.Println("Creating user collection ...")
	_, err = database.CreateCollection(nil, userCollectionName, nil)
	if err != nil {
		return errors.Wrap(err, "can't create user collection")
	}

	fmt.Println("Creating resource collection ...")
	_, err = database.CreateCollection(nil, resourcesCollectionName, nil)
	if err != nil {
		return errors.Wrap(err, "can't create resource collection")
	}

	fmt.Println("Creating graph")
	graph, err := database.CreateGraph(nil, graphName, nil)
	if err != nil {
		return errors.Wrap(err, "can't create graph")
	}

	vertexConstraints := driver.VertexConstraints{
		From: []string{userCollectionName},
		To:   []string{resourcesCollectionName},
	}

	fmt.Println("Creating hasResources edge collection")
	_, err = graph.CreateEdgeCollection(nil, EdgeHasResources, vertexConstraints)
	if err != nil {
		return errors.Wrap(err, "can't create edge collection")
	}

	return nil
}
