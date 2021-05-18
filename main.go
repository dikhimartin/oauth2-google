package main

import (
	"oauth/src"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	app.Use(middleware.Logger())
	if os.Getenv("APP_ENV") == "production" {
		app.Use(middleware.Recover())
	}

	src.RouteMaster(app.Group("/api/v1/oauth"))

	app.Logger.Fatal(app.Start(":" + os.Getenv("PORT")))
}
