package store

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	Database   string
	Collection string
	Session    *mgo.Session
}

func (r Repository) GetProducts() Products {
	// defer r.Session.Close()

	c := r.Session.DB(r.Database).C(r.Collection)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil {
		log.Fatalf("Failed to write results: %v", err)
	}

	return results
}

func (r Repository) GetProductById(id int) Product {
	defer r.Session.Close()

	c := r.Session.DB(r.Database).C(r.Collection)
	result := Product{}

	fmt.Println("ID in GetProductById", id)
	if err := c.FindId(id).One(&result); err != nil {
		log.Fatalf("Failed to write results: %v", err)
	}
	return result
}

func (r Repository) GetProductsByString(query string) Products {
	defer r.Session.Close()

	c := r.Session.DB(r.Database).C(r.Collection)
	result := Products{}

	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
		and[i] = bson.M{"title": bson.M{
			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
		}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(filter).Limit(5).All(&result); err != nil {
		log.Fatalln("Failed to write result:", err)
	}
	return result
}

func (r Repository) AddProduct(product Product) error {
	defer r.Session.Close()

	if err := r.Session.DB(r.Database).C(r.Collection).Insert(product); err != nil {
		log.Fatalf("Failed to insert product: %s", err)
	}

	return nil
}

func (r Repository) UpdateProduct(product Product) error {
	defer r.Session.Close()

	if _, err := r.Session.DB(r.Database).C(r.Collection).UpsertId(product.ID, product); err != nil {
		log.Fatalf("Failed to update product: %s", err)
	}
	return nil
}

func (r Repository) DeleteProduct(id int) error {
	defer r.Session.Close()
	if err := r.Session.DB(r.Database).C(r.Collection).RemoveId(id); err != nil {
		log.Fatalf("Failed to delete product: %s", err)
	}
	return nil
}
