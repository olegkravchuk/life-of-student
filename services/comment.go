package services

import (
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CommentService struct {
	comment models.Comment
}

func (service CommentService) GetComments(db *mgo.Database, selector bson.M) (*[]models.Comment, error) {
	comments := []models.Comment{}
	err := db.C("comments").Find(selector).Sort("-create_date").All(&comments)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}

func (service CommentService) CreateComment(db *mgo.Database, comment models.Comment) error {
	err := db.C("comments").Insert(comment)
	return err
}
