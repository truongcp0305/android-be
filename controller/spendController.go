package controller

import (
	"android-be/model"
	"android-be/service"

	"github.com/labstack/echo/v4"
)

type SpenController struct {
	spend *service.SpendService
}

func NewSpendController(s *service.SpendService) SpenController {
	return SpenController{
		spend: s,
	}
}

func (s *SpenController) ListSpend(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	sps, err := s.spend.ListSpend(params.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": sps,
	})
}

type CreateSpendParams struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id" param:"user_id"`
	Money    string `json:"money"`
	Icon     string `json:"icon"`
	Category string `json:"category"`
	Time     string `json:"time"`
	Type     string `json:"type"`
	Note     string `json:"note"`
}

func (s *SpenController) CreateSpend(c echo.Context) error {
	var params CreateSpendParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	spend := model.Spending{
		Id:       params.Id,
		UserId:   params.UserId,
		Money:    params.Money,
		Icon:     params.Icon,
		Category: params.Category,
		Time:     params.Time,
		Type:     params.Type,
		Note:     params.Note,
	}
	if err = s.spend.Create(&spend); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}

func (s *SpenController) UpdateSpend(c echo.Context) error {
	var params CreateSpendParams
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	spend := model.Spending{
		Id:       params.Id,
		UserId:   params.UserId,
		Money:    params.Money,
		Icon:     params.Icon,
		Category: params.Category,
		Time:     params.Time,
		Type:     params.Type,
		Note:     params.Note,
	}
	if err = s.spend.UpdateSpend(&spend); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}

func (s *SpenController) GetById(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	spend, err := s.spend.GetBill(params.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, spend)
}

func (s *SpenController) GetInWeek(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	spends, err := s.spend.GetInWeek(params.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": spends,
	})
}

func (s *SpenController) Delete(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	if err = s.spend.Delete(params.Id); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}
