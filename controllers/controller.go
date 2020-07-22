package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

// Index GET
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.store.GetProducts()
	data, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("failed to marhal mongo results: %s", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// CreateUser POST
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	userName := r.URL.Query().Get("user_name")
	passWord := r.URL.Query().Get("pass_word")
	roleType := mgo.Role(r.URL.Query().Get("role_type"))

	c.user.User = &mgo.User{
		Username: userName,
		Password: passWord,
		Roles:    []mgo.Role{roleType},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := c.user.AddUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data, err := json.Marshal(err)
		if err != nil {
			log.Fatalf("failed to marhal mongo error: %s", err)
		}
		w.Write(data)
	} else {
		data, _ := json.Marshal("User has been created")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
