package main

import (
	"log"

	"github.com/someshkoli/jobHandler/pkg/api"
)

func main() {
	// creating new instance of jobHandler
	jobHandler := api.MakeJobHandler()

	// Registering new mux
	mux := jobHandler.RegisterHandlers()

	// Creating new server instance with given mux
	srv := api.NewServer(mux)

	if jobHandler.StartServer(srv) != nil {
		log.Fatal("unable to start server")
	}
}
