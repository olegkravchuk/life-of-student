package models

import (
	"github.com/martini-contrib/sessionauth"
	"github.com/olegkravchuk/life_of_student/utils"
	"gopkg.in/mgo.v2/bson"
)

type MyUser struct {
	Id            int64  `form:"id" bson:"id"`
	Username      string `form:"username" bson:"username"`
	Email         string `form:"email" bson:"email"`
	Password      string `form:"password" bson:"password"`
	authenticated bool   `form:"-" db:"-"`
}

func GenerateAnonymousUser() sessionauth.User {
	return &MyUser{}
}

func (u *MyUser) Login() {
	u.authenticated = true
}

func (u *MyUser) Logout() {
	u.authenticated = false
}

func (u *MyUser) IsAuthenticated() bool {
	return u.authenticated
}

func (u *MyUser) UniqueId() interface{} {
	return u.Id
}

func (u *MyUser) GetById(id interface{}) error {
	userCollection := utils.DB.C("users")
	MyUserDoc := &MyUser{}
	err := userCollection.Find(bson.M{"id": id}).One(&MyUserDoc)
	*u = *MyUserDoc
	if err != nil {
		return err
	}

	return nil
}
