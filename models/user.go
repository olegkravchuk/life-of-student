package models

import (
	"github.com/martini-contrib/sessionauth"
	"gopkg.in/mgo.v2"
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
	var userCollection *mgo.Collection
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	userCollection = session.DB("blog").C("users")
	MyUserDoc := &MyUser{}
	er := userCollection.Find(bson.M{"id": id}).One(&MyUserDoc)
	*u = *MyUserDoc
	if er != nil {
		return err
	}

	return nil
}
