package arango

import (
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/pkg/errors"
)

type bindingVariables map[string]interface{}

var (
	arangoClient      driver.Client
	arangoDatabase    driver.Database
	arangoGraph       driver.Graph
	arangoCollections = map[string]driver.Collection{}
)

const (
	databaseName = "fleet-commander"
	graphName    = "fleet-commander-graph"
)

// Setup tries to get a arango client, database and graph. If this is not possible, it
// will return an error.
func Setup() error {
	if _, err := getClient(); err != nil {
		return err
	}

	if _, err := getDatabase(); err != nil {
		return err
	}

	return nil
}

func getClient() (driver.Client, error) {
	if arangoClient == nil {
		fmt.Println("Initializing ArangoDB client")
		endpoint := "http://localhost:8529"

		connectionConfig := http.ConnectionConfig{
			Endpoints: []string{endpoint},
		}

		connection, err := http.NewConnection(connectionConfig)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting client")
		}

		clientConfig := driver.ClientConfig{
			Connection:     connection,
			Authentication: driver.BasicAuthentication("root", "root"),
		}

		arangoClient, err = driver.NewClient(clientConfig)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting client")
		}
	}

	return arangoClient, nil
}

func getDatabase() (driver.Database, error) {
	if arangoDatabase == nil {
		fmt.Println("Initializing ArangoDB database")
		client, err := getClient()
		if err != nil {
			return nil, errors.Wrap(err, "error while getting database")
		}

		arangoDatabase, err = client.Database(nil, databaseName)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting database")
		}
	}

	return arangoDatabase, nil
}

func getCollection(name string) (driver.Collection, error) {
	if arangoCollections[name] == nil {
		fmt.Println("Initialize ArangoDB collection:", name)
		database, err := getDatabase()
		if err != nil {
			return nil, errors.Wrapf(err, "error while getting collection '%s'", name)
		}

		collection, err := database.Collection(nil, name)
		if err != nil {
			return nil, errors.Wrapf(err, "error while getting collection '%s'", name)
		}
		arangoCollections[name] = collection
	}

	return arangoCollections[name], nil
}
