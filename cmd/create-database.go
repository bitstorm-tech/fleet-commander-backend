package main

import (
    "gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
    "os"
    "fmt"
    "strings"
    "log"
)

func main() {
    client := arango.Connect()

    database, _ := client.Database(nil, arango.DatabaseName)

    if withParameter("drop") && database != nil {
        fmt.Println("Dropping existing database!")
        database.Remove(nil)
    }

    fmt.Println("Creating new database")
    database, err := client.CreateDatabase(nil, arango.DatabaseName, nil)

    if err != nil {
        log.Panic(err)
    }

    fmt.Println("Creating user collection")
    _, err = database.CreateCollection(nil, "users", nil)

    if err != nil {
        log.Panic(err)
    }
}

func withParameter(parameter string) bool {
    for i := 0; i < len(os.Args); i++ {
        if strings.ToLower(os.Args[i]) == strings.ToLower(parameter) {
            return true
        }
    }

    return false
}