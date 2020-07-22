package user

import (
	"log"

	"github.com/globalsign/mgo"
)

type Repository struct {
	Database string
	Session  *mgo.Session
	User     *mgo.User
}

func (r *Repository) AddUser() error {
	// defer r.Session.Close()
	db := r.Session.DB(r.Database)

	if err := db.UpsertUser(r.User); err != nil {
		log.Fatalf("Failed to write results: %v", err)
		return err
	}
	return nil
}
