package config

import (
	"android-be/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AppController struct {
	U controller.UserController
	S controller.SpenController
	P controller.PlanController
}

func NewRoute(e *echo.Echo, app AppController) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/user/login", func(c echo.Context) error {
		return app.U.Login(c)
	})
	e.POST("/user/register", func(c echo.Context) error {
		return app.U.Registry(c)
	})

	e.GET("", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})
}
