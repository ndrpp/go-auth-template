package controllers

import (
	"go-auth-template/utils"
	"go-auth-template/views/layout"
	"net/http"
)

func HandleHealthZ(logger *utils.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := utils.Encode[string](w, r, http.StatusOK, "Alive and well.")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
}

func HandleHome(logger *utils.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := utils.Render(r, w, layout.Base())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
}
