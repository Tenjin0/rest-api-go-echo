package main

import (
	"net/http"
	"rest-api-go/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Pre(middleware.RemoveTrailingSlash())

	routes.Init(app)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Logger.Fatal(app.Start(":1323"))
}
