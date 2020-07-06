package mongo

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func InitMongo() {
	var (
		conString  = "mongodb://root:secret@localhost"
		database   = "store"
		collection = "mobiles"
	)

	session, err := mgo.Dial(conString)

	if err != nil {
		panic(err)
	}

	type item struct {
		Id     bson.ObjectId `bson:"_id,omitempty" json:"_id"`
		Title  string        `bson:"Title" json:"Title"`
		Image  string        `bson:"Image" json:"Image"`
		Price  int           `bson:"Price" json:"Price"`
		Rating int           `bson:"Rating" json:"Rating"`
	}

	newColection := session.DB(database).C(collection)

	bulk := newColection.Bulk()

	var itemsArray []interface{}
	itemsArray = append(itemsArray,
		&item{Title: "Apple iMac Pro", Image: "http:://example.com/p1.jpg", Price: 5000, Rating: 4},
		&item{Title: "Google Pixel 2", Image: "http:://example.com/p2.jpg", Price: 2000, Rating: 5},
		&item{Title: "Apple iPhone X", Image: "http:://example.com/p3.jpg", Price: 3000, Rating: 5},
		&item{Title: "Google Chromebook", Image: "http:://example.com/p4.jpg", Price: 4000, Rating: 5},
		&item{Title: "Microsoft Holo Lens", Image: "http:://example.com/p5.jpg", Price: 1000, Rating: 4},
		&item{Title: "Samsung Galaxy S8", Image: "http:://example.com/p6.jpg", Price: 3000, Rating: 3},
	)
	bulk.Insert(itemsArray...)
	_, err = bulk.Run()
	if err != nil {
		panic(err)
	}

}
