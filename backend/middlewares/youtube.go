package middlewares

import (
	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YouTubeService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := os.Getenv("YOUTUBE_API_KEY")

			ctx := context.Background()
			service, err := youtube.NewService(ctx, option.WithAPIKey(key))

			if err != nil {
				logrus.Fatalf("Error creating new YouTube client: %v", err)
			}

			c.Set("yts", service)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
