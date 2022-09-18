package main

import (
	"github.com/himorishige/youtube-manager/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
