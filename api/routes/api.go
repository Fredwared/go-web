package routes

import (
	"api/app/controllers"
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
	{Method: GET, Endpoint: "/", Handler: controllers.Index},
}

func InitRoutes() {
	for _, route := range routes {
		http.HandleFunc(route.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			// Check request method
			if r.Method != string(route.Method) {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Path check
			// TODO: Check Pattern like /user/{object}/ /user/[a-zA-Z]
			if r.URL.Path != route.Endpoint {
				http.NotFound(w, r)
				return
			}

			route.Handler(w, r)
		})
	}
}
