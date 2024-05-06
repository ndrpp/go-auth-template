package controllers

import (
	"database/sql"
	"fmt"
	"go-auth-template/utils"
	"go-auth-template/views/pages"
	"net/http"
)

type User struct {
	Username string
	Password string
}

func RenderAuthForm(logger *utils.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := utils.Render(r, w, pages.AuthForm())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
}

func Login(logger *utils.Logger, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			r.ParseForm()
			user := &User{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
			}
			logger.Info(fmt.Sprintf("User name: %s, password: %s", user.Username, user.Password))

			//TODO:
			//1. get stored credentials
			//2. hash the received password
			//3. check match
			w.Header().Add("HX-Redirect", "/home")
		})
}
