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
          "RegisterHandler",
          "POST",
          "/register",
          RegisterHandler,
     },
	Route{
		"LoginHandler",
		"POST",
		"/login",
		LoginHandler,
	},
	Route{
		"BookHandler",
		"GET",
		"/pages/{count}/{offset}",
		BookHandler,
	},
	Route{
		"PageHandler",
		"GET",
		"/users/{userID}/pages/{pageID}",
		PageHandler,
	},
	Route{
		"UploadHandler",
		"POST",
		"/pages/new",
		UploadHandler,
	},
	Route{
		"ReplaceHandler",
		"POST",
		"/pages/replace",
		ReplaceHandler,
	},
	Route{
		"AccessCodeHandler",
		"GET",
		"/access/new",
		AccessCodeHandler,
	},

}
