package services

import (
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostService struct {
	post models.Post
}

func (service PostService) GetPost(db *mgo.Database, selector bson.M) (*models.Post, error) {
	post := models.Post{}
	err := db.C("posts").Find(selector).One(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (service PostService) GetPosts(db *mgo.Database, selector bson.M, page int, perPage int) (*[]models.Post, error) {
	posts := []models.Post{}
	err := db.C("posts").Find(selector).Sort("-id").Skip(page).Limit(perPage).All(&posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (service PostService) GetCountPosts(db *mgo.Database) (int, error) {
	count, err := db.C("posts").Find(nil).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (service PostService) CreatePost(db *mgo.Database, post models.Post) error {
	err := db.C("posts").Insert(post)
	return err
}

func (service PostService) UpdatePost(db *mgo.Database, selector bson.M, update bson.M) error {
	err := db.C("posts").Update(selector, update)
	return err
}

func (service PostService) DeletePost(db *mgo.Database, selector bson.M) error {
	err := db.C("posts").Remove(selector)
	return err
}
