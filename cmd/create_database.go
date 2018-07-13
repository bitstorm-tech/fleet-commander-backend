package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/arangodb/go-driver"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func main() {
	flag.Bool("drop", false, "drops the database before creating a new one")
	flag.Bool("help", false, "shows this help")
	flag.Parse()

	if flagPresent("help") {
		flag.Usage()
		return
	}

	client, err := arango.GetClient()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	database, _ := client.Database(nil, arango.DatabaseName)

	if flagPresent("drop") && database != nil {
		fmt.Println("Dropping existing database!")
		database.Remove(nil)
	}

	fmt.Println("Creating database:", arango.DatabaseName)
	database, err = client.CreateDatabase(nil, arango.DatabaseName, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Creating collection:", arango.CollectionUser)
	if _, err = database.CreateCollection(nil, arango.CollectionUser, nil); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Creating collection:", arango.CollectionResource)
	if _, err = database.CreateCollection(nil, arango.CollectionResource, nil); err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Creating graph:", arango.GraphName)
	graph, err := database.CreateGraph(nil, arango.GraphName, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Creating edge:", arango.EdgeHasResources)
	vertexConstraints := driver.VertexConstraints{
		From: []string{arango.CollectionUser},
		To:   []string{arango.CollectionResource},
	}

	if _, err := graph.CreateEdgeCollection(nil, arango.EdgeHasResources, vertexConstraints); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func flagPresent(_flag string) bool {
	flags := flag.Args()

	for i := 0; i < len(flags); i++ {
		if strings.ToLower(flags[i]) == strings.ToLower(_flag) {
			return true
		}
	}

	return false
}
