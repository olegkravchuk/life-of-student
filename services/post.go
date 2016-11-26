package services

import (
	"github.com/life_of_student/models"
	"github.com/life_of_student/utils"
	"gopkg.in/mgo.v2/bson"
)

type PostService struct {
	post models.Post
}

func (service PostService) GetPost(selector bson.M) (*models.Post, error) {
	post := models.Post{}
	err := utils.DB.C("posts").Find(selector).One(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (service PostService) GetPosts(selector bson.M, page int, perPage int) (*[]models.Post, error) {
	posts := []models.Post{}
	err := utils.DB.C("posts").Find(selector).Sort("-id").Skip(page).Limit(perPage).All(&posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (service PostService) GetCountPosts() (int, error) {
	count, err := utils.DB.C("posts").Find(nil).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (service PostService) CreatePost(post models.Post) error {
	err := utils.DB.C("posts").Insert(post)
	return err
}

func (service PostService) UpdatePost(selector bson.M, update bson.M) error {
	err := utils.DB.C("posts").Update(selector, update)
	return err
}

func (service PostService) DeletePost(selector bson.M) error {
	err := utils.DB.C("posts").Remove(selector)
	return err
}
