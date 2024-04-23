package main

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/ism0080/random-song-lyric-v2/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to render")
	}

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return render(c, http.StatusOK, views.Index())
	})

	e.Logger.Fatal(e.Start(":42069"))
}
