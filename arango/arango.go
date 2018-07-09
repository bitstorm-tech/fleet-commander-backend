package arango

import (
	"fmt"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/pkg/errors"
)

type bindingVariables map[string]interface{}

type edge struct {
	from string `json:"_from"`
	to   string `json:"_to"`
}

var (
	arangoClient      driver.Client
	arangoDatabase    driver.Database
	arangoGraph       driver.Graph
	arangoCollections = map[string]driver.Collection{}
)

var (
	DatabaseName       = "fleet-commander"
	GraphName          = "fleet-commander-graph"
	CollectionUser     = "user"
	CollectionResource = "resource"
	EdgeHasResources   = "hasResources"
)

// Setup tries to get a arango client, database and graph. If this is not possible, it
// will return an error.
func Setup() error {
	if _, err := GetClient(); err != nil {
		return err
	}
	if _, err := GetDatabase(); err != nil {
		return err
	}
	if _, err := GetGraph(); err != nil {
		return err
	}

	return nil
}

func GetClient() (driver.Client, error) {
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

func GetDatabase() (driver.Database, error) {
	if arangoDatabase == nil {
		fmt.Println("Initializing ArangoDB database")
		client, err := GetClient()
		if err != nil {
			return nil, errors.Wrap(err, "error while getting database")
		}

		arangoDatabase, err = client.Database(nil, DatabaseName)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting database")
		}
	}

	return arangoDatabase, nil
}

func GetGraph() (driver.Graph, error) {
	if arangoGraph == nil {
		fmt.Println("Initializing ArangoDB graph")
		database, err := GetDatabase() //.Graph(nil, GraphName)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting graph")
		}

		arangoGraph, err = database.Graph(nil, GraphName)
		if err != nil {
			return nil, errors.Wrap(err, "error while getting graph")
		}
	}

	return arangoGraph, nil
}

func getCollection(name string) (driver.Collection, error) {
	if arangoCollections[name] == nil {
		fmt.Println("Initialize ArangoDB collection:", name)
		database, err := GetDatabase()
		if err != nil {
			return nil, errors.Wrapf(err, "error while getting collection '%s'\n", name)
		}

		collection, err := database.Collection(nil, name)
		if err != nil {
			return nil, errors.Wrapf(err, "error while getting collection '%s'\n", name)
		}
		arangoCollections[name] = collection
	}

	return arangoCollections[name], nil
}

func CreateEdge(from Persistable, to Persistable, collection string) error {
	graph, err := GetGraph()
	if err != nil {
		return errors.Wrapf(err, "error while creating edge in collection '%s'\n", collection)
	}

	edges, _, err := graph.EdgeCollection(nil, collection)
	if err != nil {
		return errors.Wrapf(err, "error while creating edge in collection '%s'\n", collection)
	}
	//edges := getCollection(collection)

	fmt.Println("From: ", from)
	fmt.Println("To:   ", to)
	edge := &edge{from.ID(), to.ID()}
	fmt.Println("Edge: ", edge)

	if _, err := edges.CreateDocument(nil, edge); err != nil {
		return errors.Wrapf(err, "error while creating edge in collection '%s'\n", collection)
	}

	return nil
}

func RemoveEdge(persistable Persistable) error {
	graph, err := GetGraph()
	if err != nil {
		return errors.Wrapf(err, "error while removing edge: %v\n", persistable)
	}

	edges, _, err := graph.EdgeCollection(nil, persistable.Collection())
	if err != nil {
		return errors.Wrapf(err, "error while removing edge: %v\n", persistable)
	}

	if _, err := edges.RemoveDocument(nil, persistable.Key()); err != nil {
		return errors.Wrapf(err, "error while removing edge: %v\n", persistable)
	}

	return nil
}

func CreateDocument(persistable Persistable) error {
	collection, err := getCollection(persistable.Collection())
	if err != nil {
		return errors.Wrapf(err, "error while creating document: %v\n", persistable)
	}

	result := new(User)
	_, err = collection.CreateDocument(driver.WithReturnNew(nil, result), persistable)
	if err != nil {
		return errors.Wrapf(err, "error while creating document: %v\n", persistable)
	}

	if err != nil {
		return errors.Wrapf(err, "error while creating document: %v\n", persistable)
	}

	return nil
}

func RemoveDocument(persistable Persistable) error {
	collection, err := getCollection(persistable.Collection())
	if err != nil {
		return errors.Wrapf(err, "error while removing document: %v\n", persistable)
	}

	_, err = collection.RemoveDocument(nil, persistable.Key())
	if err != nil {
		return errors.Wrapf(err, "error while removing document: %v\n", persistable)
	}

	return nil
}
