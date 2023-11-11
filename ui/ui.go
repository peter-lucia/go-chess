package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartUI() {
	// Create a new instance of Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/css", "ui/css/")
	e.Static("/js", "ui/js/")
	e.Static("/img", "ui/img/")
	e.File("/", "ui/index.html")

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
