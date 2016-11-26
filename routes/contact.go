package routes

import (
	"encoding/base64"
	"fmt"
	"github.com/life_of_student/models"
	"github.com/life_of_student/services"
	"github.com/life_of_student/utils"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"
)

func ContactViewHandler(rnd render.Render, user sessionauth.User) {
	data := make(map[string]interface{})
	data["user"] = user.(*models.MyUser)
	rnd.HTML(200, "contact-us", data)
}

func CreateContactHandler(r *http.Request, rnd render.Render, user sessionauth.User) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	contactDoc := models.Contact{Name: name, Email: email, Subject: subject, Message: message}
	serviceContact := services.ContactService{}
	serviceContact.CreateContact(contactDoc)

	go sendEmail(contactDoc.Subject, contactDoc.Message)

	rnd.Redirect("/")
}

func sendEmail(subject, message string) {
	auth := smtp.PlainAuth("", utils.USERNAME, utils.PASSWORD, utils.HOST)

	from := mail.Address{"testFrom", utils.EMAILFROM}
	to := mail.Address{"testTo", utils.EMAILTO}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	messages := ""
	for k, v := range header {
		messages += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	messages += "\r\n" + base64.StdEncoding.EncodeToString([]byte(message))

	err := smtp.SendMail(utils.HOST+":"+utils.PORT,
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(messages))
	if err != nil {
		log.Print("ERROR: attempting to send a mail ", err)
	}
}

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}
