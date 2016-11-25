package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"

	"../models"
	"../services"
	"../utils"
)

func CreateViewPostHandler(rnd render.Render, user sessionauth.User) {
	data := make(map[string]interface{})
	data["post"] = models.Post{}
	data["user"] = user.(*models.MyUser)
	rnd.HTML(200, "create", data)
}

func CreatePostHandler(r *http.Request, rnd render.Render, database *mgo.Database, user sessionauth.User) {
	title := r.FormValue("title")
	descriptionMarkdown := r.FormValue("description")
	description := utils.ConvertMarkdownToHtml(descriptionMarkdown)

	postService := services.PostService{}
	count, _ := postService.GetCountPosts(database)

	postService.CreatePost(database, models.Post{Id: count + 1, Title: title, Description: description,
		DescriptionMarkdown: descriptionMarkdown, Author: user.(*models.MyUser), CreateDate: time.Now()})

	rnd.Redirect("/")
}

func GetHtmlMarkdownHandler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	outputHtml := utils.ConvertMarkdownToHtml(md)

	rnd.JSON(200, map[string]interface{}{"html": outputHtml})
}

func EditViewPostHandler(w http.ResponseWriter, r *http.Request, rnd render.Render, params martini.Params,
	database *mgo.Database, user sessionauth.User) {
	id, _ := params["id"]
	idInt, _ := strconv.Atoi(id)

	servicePost := services.PostService{}
	post, err := servicePost.GetPost(database, bson.M{"id": idInt})
	if err != nil {
		http.NotFound(w, r)
	}
	data := make(map[string]interface{})
	data["post"] = post
	data["user"] = user.(*models.MyUser)
	rnd.HTML(200, "edit", data)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request, rnd render.Render, params martini.Params,
	database *mgo.Database, user sessionauth.User) {
	id, _ := params["id"]
	idInt, _ := strconv.Atoi(id)
	title := r.FormValue("title")
	descriptionMarkdown := r.FormValue("description")
	description := utils.ConvertMarkdownToHtml(descriptionMarkdown)

	servicePost := services.PostService{}
	err := servicePost.UpdatePost(
		database,
		bson.M{"id": idInt},
		bson.M{"$set": bson.M{"title": title, "description": description, "descriptionmarkdown": descriptionMarkdown}},
	)
	if err != nil {
		http.NotFound(w, r)
	}
	rnd.Redirect("/")
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request, params martini.Params, database *mgo.Database) {
	id := params["id"]
	idInt, _ := strconv.Atoi(id)
	servicePost := services.PostService{}
	err := servicePost.DeletePost(database, bson.M{"id": idInt})
	if err != nil {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, "/", 302)
}

func ViewPostHandler(w http.ResponseWriter, r *http.Request, rnd render.Render, params martini.Params, database *mgo.Database, user sessionauth.User) {
	id, _ := params["id"]
	idInt, _ := strconv.Atoi(id)

	postService := services.PostService{}
	post, err := postService.GetPost(database, bson.M{"id": idInt})
	if err != nil {
		http.NotFound(w, r)
	}

	serviceComment := services.CommentService{}
	comments, _ := serviceComment.GetComments(database, bson.M{"post": post})

	data := make(map[string]interface{})
	data["post"] = post
	data["comments"] = comments
	data["user"] = user.(*models.MyUser)

	rnd.HTML(200, "view-post", data)

}
