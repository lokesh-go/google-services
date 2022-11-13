package main

import (
	"log"

	"github.com/lokesh-go/google-services/src/initialise"
)

func main() {
	// Initialises
	err := initialise.Init()
	if err != nil {
		log.Fatal("failed to initialises: ", err)
	}
}
