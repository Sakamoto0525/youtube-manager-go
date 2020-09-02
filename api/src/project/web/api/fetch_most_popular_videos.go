package api

import (
	"context"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPoplarVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := "AIzaSyCJNFuzmAGeAe5ZCXSqgo-el8ysXmnxxMQ"

		ctx := context.Background()

		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))
		if err != nil {
			logrus.Fatalf("Error createing new Youtube service: %v", err)
		}

		call := yts.Videos.List([]string{"id", "snippet"}).Chart("mostPopular").MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
