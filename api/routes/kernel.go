package routes

import (
	"api/routes/middlewares"
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Route struct {
	Method     Method
	Endpoint   string
	Handler    http.HandlerFunc
	middleware Middleware
}

var DefaultMiddleware []Middleware

func UseMiddleWares(handler http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}

func Kernel() {
	DefaultMiddleware = []Middleware{
		middlewares.JSONMiddleware,
	}

	for _, route := range routes {
		http.HandleFunc(route.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			UseMiddleWares(route.Handler, DefaultMiddleware...)

			// Check request method
			if r.Method != string(route.Method) {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Path check
			// TODO: Check Pattern like /user/{object}/ /user/[a-zA-Z]
			if r.URL.Path != route.Endpoint {
				response := map[string]string{"error": "Not Found"}
				jsonResponse, err := json.Marshal(response)

				// Server error
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusNotFound)
				_, err = w.Write(jsonResponse)
				if err != nil {
					return
				}
				return
			}

			route.Handler(w, r)
		})
	}
}
