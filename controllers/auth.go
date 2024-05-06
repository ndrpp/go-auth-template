package controllers

import (
	"database/sql"
	"fmt"
	"go-auth-template/utils"
	"go-auth-template/views/pages"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	Username string
	Password string
}

func RenderAuthForm(logger *utils.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var action string
			if r.URL.String() == "/register" {
				action = "REGISTER"
			} else {
				action = "LOGIN"
			}
			err := utils.Render(r, w, pages.AuthForm(action))
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
			//TODO
			//perform some validation on the data
			user := &User{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
			}
			logger.Info(fmt.Sprintf("User name: %s, password: %s", user.Username, user.Password))

			//TODO:
			//1. get stored credentials
			//2. hash the received password
			//3. check match
			//4. store JWT on FE if match
			w.Header().Add("HX-Redirect", "/home")
		})
}

func Register(logger *utils.Logger, db *sql.DB) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			//TODO
			//perform some validation on the data
			//then hash the received password
			new_user := &User{
				Username: r.FormValue("username"),
				Password: r.FormValue("password"),
			}

			_, err := db.Exec("INSERT INTO users (user_id, username, password) VALUES ($1, $2, $3)", uuid.NewString(), new_user.Username, new_user.Password)
			if err != nil {
				logger.Error(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//store JWT on FE
			w.Header().Add("HX-Redirect", "/home")
		})
}
