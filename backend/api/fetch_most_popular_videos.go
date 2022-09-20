package api

import (
	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		key := os.Getenv("YOUTUBE_API_KEY")

		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))

		if err != nil {
			logrus.Fatalf("Error creating new YouTube client: %v", err)
		}

		k := []string{"id", "snippet"}
		call := yts.Videos.List(k).Chart("mostPopular").MaxResults(3)

		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error making API call to YouTube: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
