package main

import (
	"fmt"
)

func main() {
	fmt.Println("Server listening on port 8080...")

	// Initialize kernel with routes
	UseApi()

	// Initialize server
	UseServer()
}
