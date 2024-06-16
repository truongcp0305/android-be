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

type UserIdPath struct {
	Id string `param:"id"`
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

type RegistryParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserController) Registry(c echo.Context) error {
	var params RegistryParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	user := model.User{
		Username: params.Username,
		Password: params.Password,
	}
	uid, err := u.user.CreateUser(&user)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"id": uid,
	})
}

func (u *UserController) GetInfo(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	user, err := u.user.GetInfo(params.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, user)
}

func (u *UserController) AdLogin(c echo.Context) error {
	var params LoginParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	ad := model.AdModel{
		Username: params.Username,
		Password: params.Password,
	}
	if err = u.user.AdLogin(&ad); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"id": ad.Id,
	})
}

func (u *UserController) ChangeInfo(c echo.Context) error {
	var params model.User
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	if err := u.user.ChangeUserInfo(&params); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"message": "Success",
	})
}

type SearchParams struct {
	Data string `json:"data" query:"data"`
}

func (u *UserController) Search(c echo.Context) error {
	var params SearchParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	us, err := u.user.Search(params.Data)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"Data": us,
	})
}

type QueryUserParam struct {
	Page int `query:"page"`
}

func (u *UserController) QueryUser(c echo.Context) error {
	var params QueryUserParam
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	us, err := u.user.QueryUser(params.Page)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"Data": us,
	})
}
