package routes

import (
	"project/web/api"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPoplarVideos())  // Video一覧取得
		g.GET("/video/:id", api.GetVideo())             // Video詳細取得
		g.GET("/related/:id", api.FetchRelatedVideos()) // 関連動画取得
	}
}
