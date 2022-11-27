package main

import (
	"log"

	"github.com/lokesh-go/google-services/src/example"
)

func main() {
	// Initialises
	err := example.Init()
	if err != nil {
		log.Fatal("failed to initialises: ", err)
	}
}
