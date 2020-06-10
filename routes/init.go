package routes

import "github.com/labstack/echo/v4"

func Init(app *echo.Echo) {

	api := app.Group("/api")
	RegisterUser(api)
}
