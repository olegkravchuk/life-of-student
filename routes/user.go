package routes

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/olegkravchuk/life_of_student/models"
	"github.com/olegkravchuk/life_of_student/services"
	"github.com/olegkravchuk/life_of_student/utils"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func RegistrationViewHandler(rnd render.Render, user sessionauth.User) {
	rnd.HTML(200, "registration", nil)
}

func RegistrationUserHandler(r *http.Request, rnd render.Render, session sessions.Session) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	serviceUsers := services.UserService{}
	n, _ := serviceUsers.GetCountUsers()
	newUser := models.MyUser{Id: int64(n + 1), Username: username, Email: email, Password: utils.PasswordToHash(password)}
	serviceUsers.CreateUser(newUser)
	//TODO: need add unique field, for example email

	err := sessionauth.AuthenticateSession(session, &newUser)
	if err != nil {
		rnd.JSON(500, err)
	}

	rnd.Redirect("/")
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request, rnd render.Render, session sessions.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	serviceUser := services.UserService{}
	getUser, err := serviceUser.GetUser(bson.M{"username": username})
	if err != nil || !utils.CompareHashAndPassword(getUser.Password, password) {
		rnd.Redirect(sessionauth.RedirectUrl)
		return
	} else {
		err := sessionauth.AuthenticateSession(session, getUser)
		if err != nil {
			rnd.JSON(500, err)
		}

		rnd.Redirect("/")
		return
	}

}

func LogoutUserHandler(session sessions.Session, user sessionauth.User, rnd render.Render) {
	sessionauth.Logout(session, user)
	rnd.Redirect("/")
}
