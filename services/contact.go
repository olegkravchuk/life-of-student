package services

import (
	"../models"
	"gopkg.in/mgo.v2"
)

type ContactService struct {
	contact models.Contact
}

func (service ContactService) CreateContact(db *mgo.Database, contact models.Contact) error {
	err := db.C("contact").Insert(contact)
	return err
}
