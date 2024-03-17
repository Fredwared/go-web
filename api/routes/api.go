package routes

import (
	"fmt"
	"net/http"
)

type Method string

const (
	GET     Method = "GET"
	POST           = "POST"
	PUT            = "PUT"
	DELETE         = "DELETE"
	PATCH          = "PUT"
	OPTIONS        = "OPTIONS"
)

type Route struct {
	Method   Method
	Endpoint string
	Handler  http.HandlerFunc
}

var routes = []Route{
	{Method: POST, Endpoint: "/", Handler: IndexController},
	{Method: POST, Endpoint: "/Hello", Handler: OtherController},
}

func InitRoutes() {
	for _, route := range routes {
		http.HandleFunc(route.string, route.Handler)
	}
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func OtherController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WWWWWW")
}
