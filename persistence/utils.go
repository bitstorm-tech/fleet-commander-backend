package persistence

import (
	"fmt"
	"os"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func newArangoClient() driver.Client {
	endpoint := "http://localhost:8529"

	connectionConfig := http.ConnectionConfig{
		Endpoints: []string{endpoint},
	}

	connection, err := http.NewConnection(connectionConfig)
	if err != nil {
		fmt.Println("ERROR: Can't create connection", err)
		os.Exit(1)
	}

	clientConfig := driver.ClientConfig{
		Connection:     connection,
		Authentication: driver.BasicAuthentication("root", "root"),
	}

	c, err := driver.NewClient(clientConfig)
	if err != nil {
		fmt.Println("ERROR: Can't create ArangoDB client", err)
		os.Exit(1)
	}

	return c
}
