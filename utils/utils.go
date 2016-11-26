package utils

import (
	"github.com/russross/blackfriday"
	"golang.org/x/crypto/bcrypt"
	"html/template"
)

func Unescape(s string) interface{} {
	return template.HTML(s)
}

func SliceStr(s string) string {
	if len(s) > 100 {
		return s[:100] + "..."
	} else {
		return s
	}
}

func ConvertMarkdownToHtml(str string) string {
	return string(blackfriday.MarkdownBasic([]byte(str)))
}

func PasswordToHash(s string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func CompareHashAndPassword(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
