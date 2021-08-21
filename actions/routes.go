package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// RoutesHandler is Buffalo handler to show routes
func RoutesHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("routes.html"))
}
