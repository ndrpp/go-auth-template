package utils

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(c context.Context, w http.ResponseWriter, component templ.Component) error {
	return component.Render(c, w)
}
