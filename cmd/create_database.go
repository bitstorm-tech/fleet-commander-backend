package main

import (
	"flag"
	"fmt"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func main() {
	drop := flag.Bool("drop", false, "drops the database before creating a new one")
	help := flag.Bool("help", false, "shows this help")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	err := arango.SetupFleetCommanderDatabase(*drop)
	if err != nil {
		fmt.Printf("--- ERROR ---\n%+v", err)
	} else {
		fmt.Println("--- DONE ---")
	}
}
