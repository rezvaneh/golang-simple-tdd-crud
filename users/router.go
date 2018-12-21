package users

import (
	"github.com/gorilla/mux"
	"net/http"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"Insert",
		"POST",
		"/Insert",
		controller.InsertUser,
	},
	Route{
		"New",
		"GEt",
		"/New",
		controller.New,
	},
	Route{
		"Edit",
		"GET",
		"/Edit/{id}",
		controller.GetUserForEdit,
	},
	Route{
		"Update",
		"POST",
		"/Update",
		controller.UpdateUser,
	},
	Route{
		"Show",
		"GET",
		"/Show/{id}",
		controller.GetUser,
	},
	Route{
		"Delete",
		"GET",
		"/Delete/{id}",
		controller.DeleteUser,
	},
	Route{
		"Search",
		"POST",
		"/Search",
		controller.SearchUser,
	}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
		http.Handle(route.Pattern, router)
	}

	return router
}
