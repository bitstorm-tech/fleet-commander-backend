package arango

import (
    "fmt"
    "os"

    "github.com/arangodb/go-driver"
    "github.com/arangodb/go-driver/http"
)

type bindingVariables map[string]interface{}

var (
    arangoClient      driver.Client
    arangoDatabase    driver.Database
    arangoCollections map[string]driver.Collection
)

var (
    DatabaseName = "fleet-commander"
)

// Connect tries to get a arango client. If this is not possible, it
// will will exit. If the arango client is already crated, nothing
// happens.
func Connect() driver.Client {
    if arangoClient == nil {
        arangoClient = getClient()
    }

    return arangoClient
}

func getClient() driver.Client {
    if arangoClient == nil {
        fmt.Println("Initializing ArangoDB client")
        endpoint := "http://localhost:8529"

        connectionConfig := http.ConnectionConfig{
            Endpoints: []string{endpoint},
        }

        connection, err := http.NewConnection(connectionConfig)
        if err != nil {
            fmt.Println("ERROR: can't create connection:", err)
            os.Exit(1)
        }

        clientConfig := driver.ClientConfig{
            Connection:     connection,
            Authentication: driver.BasicAuthentication("root", "root"),
        }

        arangoClient, err = driver.NewClient(clientConfig)
        if err != nil {
            fmt.Println("ERROR: can't create ArangoDB client:", err)
            os.Exit(1)
        }
    }

    return arangoClient
}

func getDatabase() driver.Database {
    if arangoDatabase == nil {
        fmt.Println("Initializing ArangoDB database")
        client := getClient()
        var err error
        arangoDatabase, err = client.Database(nil, DatabaseName)
        if err != nil {
            fmt.Println("ERROR: can't get database:", err)
            os.Exit(1)
        }
    }

    return arangoDatabase
}

func getCollection(name string) driver.Collection {
    if arangoCollections[name] == nil {
        fmt.Println("Initialize ArangoDB collection:", name)
        database := getDatabase()
        collection, err := database.Collection(nil, name)
        if err != nil {
            fmt.Println("ERROR: can't get collection", name, err)
            os.Exit(1)
        }
        arangoCollections[name] = collection
    }

    return arangoCollections[name]
}
