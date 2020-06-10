package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(api *echo.Group) {

	defaultRoute := func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}

	users := api.Group("/users")
	users.POST("", defaultRoute)
	users.GET("", defaultRoute)
	users.GET("/:id", defaultRoute)
	users.PUT("/:id", defaultRoute)
	users.DELETE("/:id", defaultRoute)
}
