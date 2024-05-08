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

	e.POST("/spend", func(c echo.Context) error {
		return app.S.CreateSpend(c)
	})
	e.GET("/spends/:id", func(c echo.Context) error {
		return app.S.ListSpend(c)
	})
	e.GET("/spend/:id", func(c echo.Context) error {
		return app.S.GetById(c)
	})
	e.PUT("/spend", func(c echo.Context) error {
		return app.S.UpdateSpend(c)
	})
	e.DELETE("/spend/:id", func(c echo.Context) error {
		return app.S.Delete(c)
	})
	e.GET("/spend/in-week/:id", func(c echo.Context) error {
		return app.S.GetInWeek(c)
	})

	e.POST("/plan", func(c echo.Context) error {
		return app.P.Create(c)
	})
	e.GET("/plans/:id", func(c echo.Context) error {
		return app.P.ListPlan(c)
	})
	e.PUT("/plan", func(c echo.Context) error {
		return app.P.Update(c)
	})
	e.DELETE("/plan/:id", func(c echo.Context) error {
		return app.P.Delete(c)
	})

	e.GET("", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})
}
