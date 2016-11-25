package models

import (
	"time"
)

type Post struct {
	Id                  int    `bson:"id"`
	Title               string `form:"title" binding:"required"`
	Description         string `form:"description" binding:"required"`
	DescriptionMarkdown string
	Author              *MyUser   `bson:"author"`
	CreateDate          time.Time `bson:"create_date"`
}
