package controllers

import (
	"github.com/ArchieSpinos/eCommerce/store"
	"github.com/ArchieSpinos/eCommerce/user"
)

type Controller struct {
	store store.Repository
	user  user.Repository
}
