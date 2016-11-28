package services

import (
	"github.com/olegkravchuk/life_of_student/models"
	"github.com/olegkravchuk/life_of_student/utils"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	user models.MyUser
}

func (service UserService) CreateUser(user models.MyUser) error {
	err := utils.DB.C("users").Insert(user)
	return err
}

func (service UserService) GetCountUsers() (int, error) {
	count, err := utils.DB.C("users").Find(nil).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (service UserService) GetUser(selector bson.M) (*models.MyUser, error) {
	user := models.MyUser{}
	err := utils.DB.C("users").Find(selector).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
