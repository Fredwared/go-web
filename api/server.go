package main

import (
	"api/routes"
	"log"
	"net/http"
)

var hostName = "localhost"
var serverPort = "8080"

func UseApi() {
	routes.Kernel()
}

func UseServer() {
	err := http.ListenAndServe(hostName+":"+serverPort, nil)

	if err != nil {
		log.Fatal(err)
	}
}
