package main

import (
	"context"
	"embed"
	"net/http"

	"github.com/a-h/templ"
	geniusapi "github.com/ism0080/random-song-lyric-v2/internal/genius-api"
	"github.com/ism0080/random-song-lyric-v2/models"
	"github.com/ism0080/random-song-lyric-v2/views"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed static/*
var staticAssets embed.FS

func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to render")
	}

	return nil
}

func main() {
	godotenv.Load()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "static",
		Filesystem: http.FS(staticAssets),
	}))

	page := models.NewPage()

	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Index(page))
	})

	e.POST("/randomLyric", func(c echo.Context) error {
		artist := c.FormValue("artist")

		randomLyric := geniusapi.GetRandomLyrics(artist)

		return render(c, http.StatusOK, views.DisplayComponent(randomLyric))
	})

	e.Logger.Fatal(e.Start(":42069"))
}
