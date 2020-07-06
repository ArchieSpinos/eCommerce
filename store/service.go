package store

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

type Repository struct{}

const (
	MondgoConString = "mongodb://root:secret@localhost"
	Database        = "store"
	Collection      = "mobiles"
)

func (r Repository) GetProducts() Products {
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()

	c := session.DB(Database).C(Collection)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil {
		log.Fatalf("Failed to write results: %v", err)
	}

	return results
}

func (r Repository) GetProductById(id int) Product {
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()

	c := session.DB(Database).C(Collection)
	result := Product{}

	fmt.Println("ID in GetProductById", id)
	if err := c.FindId(id).One(&result); err != nil {
		log.Fatalf("Failed to write results: %v", err)
	}
	return result
}

func (r Repository) GetProductsByString(query string) Products {
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()

	c := session.DB(Database).C(Collection)
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
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()

	if err = session.DB(Database).C(Collection).Insert(product); err != nil {
		log.Fatalf("Failed to insert product: %s", err)
	}

	return nil
}

func (r Repository) UpdateProduct(product Product) error {
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()

	if _, err = session.DB(Database).C(Collection).UpsertId(product.ID, product); err != nil {
		log.Fatalf("Failed to update product: %s", err)
	}
	return nil
}

func (r Repository) DeleteProduct(id int) error {
	session, err := mgo.Dial(MondgoConString)

	if err != nil {
		log.Fatalf("Failed to establish connection to Mongo server: %v", err)
	}
	defer session.Close()
	if err = session.DB(Database).C(Collection).RemoveId(id); err != nil {
		log.Fatalf("Failed to delete product: %s", err)
	}
	return nil
}
