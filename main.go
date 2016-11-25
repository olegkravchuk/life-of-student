package main

import (
	"github.com/beatrichartz/martini-sockets"
	"github.com/go-martini/martini" //for routing
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render" //Martini middleware/handler for easily rendering serialized JSON, XML, and HTML template responses
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"gopkg.in/mgo.v2"
	"html/template"

	"./models"
	"./routes"
	"./utils"
)

var postComment *routes.PostsComment

func main() {
	m := martini.Classic()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	db := session.DB("blog")
	m.Map(db)
	postComment = routes.NewPostsComment()
	m.Map(postComment)

	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(models.GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/auth/registration"
	sessionauth.RedirectParam = "/"

	unescapeFuncMap := template.FuncMap{"unescape": utils.Unescape}
	sliceStringFuncMap := template.FuncMap{"sliceStr": utils.SliceStr}

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "master",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{unescapeFuncMap, sliceStringFuncMap},
		Charset:    "UTF-8",
		IndentJSON: true,
		IndentXML:  true,
	}))

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Get("/", routes.IndexHandler)
	m.Group("/posts", func(r martini.Router) {
		r.Get("/create", sessionauth.LoginRequired, routes.CreateViewPostHandler)
		r.Post("/create", sessionauth.LoginRequired, binding.Bind(models.Post{}), routes.CreatePostHandler)
		r.Get("/edit/:id", sessionauth.LoginRequired, routes.EditViewPostHandler)
		r.Post("/edit/:id", sessionauth.LoginRequired, binding.Bind(models.Post{}), routes.EditPostHandler)
		r.Get("/delete/:id", routes.DeletePostHandler)
		r.Get("/:id", routes.ViewPostHandler)
	})
	m.Group("/auth", func(r martini.Router) {
		r.Get("/registration", routes.RegistrationViewHandler)
		r.Post("/registration", binding.Bind(models.MyUser{}), routes.RegistrationUserHandler)
		r.Post("/login", binding.Bind(models.MyUser{}), routes.LoginUserHandler)
		r.Get("/logout", sessionauth.LoginRequired, routes.LogoutUserHandler)
	})
	m.Group("/comment", func(r martini.Router) {
		r.Post("/create", routes.CreateCommentHandler)
	})
	m.Group("/contact-us", func(r martini.Router) {
		m.Get("", routes.ContactViewHandler)
		r.Post("/create", routes.CreateContactHandler)
	})

	m.Post("/markdown-html", routes.GetHtmlMarkdownHandler)
	m.Get("/sockets/posts/:id", sockets.JSON(models.Message{}), routes.CreateCommentHandler)

	m.Run()
}
