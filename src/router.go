package src

import (
	"github.com/labstack/echo"
	c "oauth/src/controller"
)

func RouteMaster(app *echo.Group) {
	app.GET("/google", c.Oauth)

}
