package routes

import (
	"api/app/controllers"
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

var routes = []Route{
	{Method: GET, Endpoint: "/", Handler: controllers.Home},
}
