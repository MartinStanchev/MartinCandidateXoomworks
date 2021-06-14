package main

import (
	"MartinCandidate/controller"
	"log"
	"net/http"
	"os"
)

func main() {
	// Initialize the controller.
	controller.ControllerInit()

	// Default to port 8080.
	port := "8080"

	for k, v := range os.Args[1:] {
		if v == "--port" {
			port = os.Args[k+2]
		}
	}

	log.Println("Running on port: " + port)
	// Start the server.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
