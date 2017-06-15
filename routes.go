package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Open1",
		"GET",
		"/api/valve/1/open",
		Open1,
	},
	Route{
		"Close1",
		"GET",
		"/api/valve/1/close",
		Close1,
	},
	Route{
		"Open2",
		"GET",
		"/api/valve/2/open",
		Open2,
	},
	Route{
		"Close2",
		"GET",
		"/api/valve/2/close",
		Close2,
	},
	Route{
		"Open3",
		"GET",
		"/api/valve/3/open",
		Open3,
	},
	Route{
		"Close3",
		"GET",
		"/api/valve/3/close",
		Close3,
	},
	Route{
		"Open4",
		"GET",
		"/api/valve/4/open",
		Open4,
	},
	Route{
		"Close4",
		"GET",
		"/api/valve/4/close",
		Close4,
	},
	Route{
		"Open1Timer",
		"GET",
		"/api/valve/1/timer",
		Open1Timer,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
