package routes

import (
	"github.com/Unknwon/paginater"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/olegkravchuk/life_of_student/models"
	"github.com/olegkravchuk/life_of_student/services"
	"net/http"
	"strconv"
)

const PERPAGE = 5

func IndexHandler(rnd render.Render, user sessionauth.User, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage := 1
	if page != "" {
		currentPage, _ = strconv.Atoi(page)
	}

	postService := services.PostService{}
	posts, _ := postService.GetPosts(nil, PERPAGE*(currentPage-1), PERPAGE)
	countPosts, _ := postService.GetCountPosts()

	data := make(map[string]interface{})
	data["posts"] = posts

	data["page"] = paginater.New(countPosts, PERPAGE, currentPage, 4)
	data["user"] = user.(*models.MyUser)
	rnd.HTML(200, "index", data)
}
