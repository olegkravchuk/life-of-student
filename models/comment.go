package models

import "time"

type Comment struct {
	Post       Post      `bson:"post" json:"post"`
	Author     MyUser    `bson:"author" json:"author"`
	Comment    string    `bson:"comment" json:"comment"`
	CreateDate time.Time `bson:"create_date" json:"create_date"`
}

type Message struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Author  string `json:"author"`
}
