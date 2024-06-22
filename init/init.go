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
	spendSv := service.NewSpendService(repo)
	planSv := service.NewPalnService(repo)

	userC := controller.NewUserController(userSv)
	spendC := controller.NewSpendController(spendSv)
	planC := controller.NewPalnController(planSv)

	ac := config.AppController{
		U: userC,
		S: spendC,
		P: planC,
	}

	e := echo.New()

	config.NewRoute(e, ac)

	go e.Logger.Fatal(e.Start("192.168.1.14:1234"))
}
