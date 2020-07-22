package controllers

import (
	"net/http"

	"github.com/ArchieSpinos/eCommerce/store"
	"github.com/ArchieSpinos/eCommerce/user"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func initMongo() *mgo.Session {
	mondgoConString := "mongodb://root:secret@localhost"
	session, err := mgo.Dial(mondgoConString)
	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	return session
}

var controler = &Controller{
	store: store.Repository{
		Database:   "store",
		Collection: "mobiles",
		Session:    initMongo(),
	},
	user: user.Repository{
		Database: "store",
		Session:  initMongo(),
		User:     &mgo.User{},
	},
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		controler.Index,
	},
	{
		"CreateUser",
		"POST",
		"/users",
		controler.CreateUser,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
	return router
}
