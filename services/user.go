package services

import (
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	user models.MyUser
}

func (service UserService) CreateUser(db *mgo.Database, user models.MyUser) error {
	err := db.C("users").Insert(user)
	return err
}

func (service UserService) GetCountUsers(db *mgo.Database) (int, error) {
	count, err := db.C("users").Find(nil).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (service UserService) GetUser(db *mgo.Database, selector bson.M) (*models.MyUser, error) {
	user := models.MyUser{}
	err := db.C("users").Find(selector).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
