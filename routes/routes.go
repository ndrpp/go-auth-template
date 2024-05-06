package routes

import (
	"database/sql"
	"go-auth-template/controllers"
	"go-auth-template/utils"
	"net/http"
)

func AddBaseRoutes(mux *http.ServeMux, logger *utils.Logger) {
	mux.Handle("/healthz", controllers.HandleHealthZ(logger))
	mux.Handle("/home", controllers.HandleHome(logger))
}

func AddAuthRoutes(mux *http.ServeMux, logger *utils.Logger, db *sql.DB) {
	mux.Handle("GET /login", controllers.RenderAuthForm(logger))
	mux.Handle("POST /login", controllers.Login(logger, db))

	mux.Handle("GET /register", controllers.RenderAuthForm(logger))
}
