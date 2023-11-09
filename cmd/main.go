package main

import (
	"log"

	"github.com/ECOMMERCE_PROJECT/pkg/config"
	"github.com/ECOMMERCE_PROJECT/pkg/di"
	
	// Add this import
)


func main() {
    config, configErr := config.LoadConfig()
    if configErr != nil {
        log.Fatal("Error in configuration", configErr)
    }
   

    server, diErr := di.InitializeAPI(config)
    if diErr != nil {
        log.Fatal("Cannot initialize API", diErr)
    }

    if server != nil {
        server.Start()
    } else {
        log.Fatal("Server is nil")
    }
}

