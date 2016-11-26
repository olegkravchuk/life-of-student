package utils

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var DB *mgo.Database

func init() {
	var err error

	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	DB = session.DB("blog")
}
