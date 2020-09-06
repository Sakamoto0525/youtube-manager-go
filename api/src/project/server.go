package main

import (
	"project/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	//Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	routes.Init(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
