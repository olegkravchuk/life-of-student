package services

import (
	"github.com/life_of_student/models"
	"github.com/life_of_student/utils"
	"gopkg.in/mgo.v2/bson"
)

type CommentService struct {
	comment models.Comment
}

func (service CommentService) GetComments(selector bson.M) (*[]models.Comment, error) {
	comments := []models.Comment{}
	err := utils.DB.C("comments").Find(selector).Sort("-create_date").All(&comments)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}

func (service CommentService) CreateComment(comment models.Comment) error {
	err := utils.DB.C("comments").Insert(comment)
	return err
}
