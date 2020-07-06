package store

import "github.com/globalsign/mgo/bson"

type Product struct {
	ID     bson.ObjectId `bson:"_id" json:"_id"`
	Title  string        `bson:"Title" json:"Title"`
	Image  string        `bson:"Image" json:"Image"`
	Price  int           `bson:"Price" json:"Price"`
	Rating int           `bson:"Rating" json:"Rating"`
}

type Products []Product
