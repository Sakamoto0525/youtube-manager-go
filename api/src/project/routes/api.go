package routes

import (
	"project/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPoplarVideos())
		g.GET("/video/:id", api.GetVideo())
	}
}
