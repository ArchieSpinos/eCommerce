package store

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Controller struct {
	Repository Repository
}

// Index GET
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetProducts()
	data, err := json.Marshal(products)
	if err != nil {
		log.Fatalf("failed to marhal mongo results: %s", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
