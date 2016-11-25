package routes

import (
	"../models"
	"../services"
	"github.com/Unknwon/paginater"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"gopkg.in/mgo.v2"
	"net/http"
	"strconv"
)

const PERPAGE = 5

func IndexHandler(rnd render.Render, database *mgo.Database, user sessionauth.User, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage := 1
	if page != "" {
		currentPage, _ = strconv.Atoi(page)
	}

	postService := services.PostService{}
	posts, _ := postService.GetPosts(database, nil, PERPAGE*(currentPage-1), PERPAGE)
	countPosts, _ := postService.GetCountPosts(database)

	data := make(map[string]interface{})
	data["posts"] = posts

	data["page"] = paginater.New(countPosts, PERPAGE, currentPage, 4)
	data["user"] = user.(*models.MyUser)
	rnd.HTML(200, "index", data)
}
