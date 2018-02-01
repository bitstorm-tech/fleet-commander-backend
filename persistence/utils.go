package persistence

import (
	"fmt"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type bindingVariables map[string]interface{}

func newArangoClient() driver.Client {
	endpoint := "http://localhost:8529"

	connectionConfig := http.ConnectionConfig{
		Endpoints: []string{endpoint},
	}

	connection, err := http.NewConnection(connectionConfig)
	if err != nil {
		fmt.Println("ERROR: can't create connection", err)
		os.Exit(1)
	}

	clientConfig := driver.ClientConfig{
		Connection:     connection,
		Authentication: driver.BasicAuthentication("root", "root"),
	}

	client, err := driver.NewClient(clientConfig)
	if err != nil {
		fmt.Println("ERROR: can't create ArangoDB client", err)
		os.Exit(1)
	}

	return client
}

func newArangoDatabase() driver.Database {
	client := newArangoClient()
	database, err := client.Database(nil, "fleet-commander")
	if err != nil {
		fmt.Println("ERROR: can't get database", err)
		os.Exit(1)
	}

	return database
}

func newArangoCollection(name string) driver.Collection {
	database := newArangoDatabase()
	collection, err := database.Collection(nil, name)
	if err != nil {
		fmt.Println("ERROR: can't get collection", name, err)
		os.Exit(1)
	}

	return collection
}
