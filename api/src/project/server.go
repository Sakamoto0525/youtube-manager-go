package main

import (
	"project/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//Middlewares
	e.Use(middleware.Logger())

	// Routes
	routes.Init(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
