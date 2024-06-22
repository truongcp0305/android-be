package config

import (
	"android-be/controller"
	"net/http"

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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/user/login", func(c echo.Context) error {
		return app.U.Login(c)
	})
	e.POST("/user/register", func(c echo.Context) error {
		return app.U.Registry(c)
	})
	e.GET("/user/:id", func(c echo.Context) error {
		return app.U.GetInfo(c)
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
	e.GET("/plan/:id", func(c echo.Context) error {
		return app.P.ListPlan(c)
	})
	e.PUT("/plan", func(c echo.Context) error {
		return app.P.Update(c)
	})
	e.DELETE("/plan", func(c echo.Context) error {
		return app.P.Delete(c)
	})

	e.GET("/plan/key", func(c echo.Context) error {
		return app.P.GetByKey(c)
	})

	e.POST("/admin/login", func(c echo.Context) error {
		return app.U.AdLogin(c)
	})
	e.PUT("/admin/user", func(c echo.Context) error {
		return app.U.ChangeInfo(c)
	})

	e.GET("/admin/search", func(c echo.Context) error {
		return app.U.Search(c)
	})

	e.GET("/user/query", func(c echo.Context) error {
		return app.U.QueryUser(c)
	})

	e.GET("", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})
}
