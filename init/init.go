package init

import (
	"android-be/config"
	"android-be/controller"
	"android-be/repository"
	"android-be/service"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	db := repository.Connn()
	repo := repository.NewDatabase(db.Collection("spend"), db.Collection("plan"), db.Collection("user"))

	userSv := service.NewUserService(repo)

	userC := controller.NewUserController(userSv)

	ac := config.AppController{
		U: userC,
	}

	e := echo.New()

	config.NewRoute(e, ac)

	go e.Logger.Fatal(e.Start(":8080"))
}
