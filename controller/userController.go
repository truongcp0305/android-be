package controller

import (
	"android-be/model"
	"android-be/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	user *service.UserService
}

func NewUserController(u *service.UserService) UserController {
	return UserController{
		user: u,
	}
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) Login(c echo.Context) error {
	var params LoginParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	userI := model.User{
		Username: params.Username,
		Password: params.Password,
	}

	user, err := u.user.Login(&userI)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"id": user.Id,
	})
}
