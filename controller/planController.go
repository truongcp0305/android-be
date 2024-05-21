package controller

import (
	"android-be/model"
	"android-be/service"

	"github.com/labstack/echo/v4"
)

type PlanController struct {
	plan *service.PlanService
}

func NewPalnController(p *service.PlanService) PlanController {
	return PlanController{
		plan: p,
	}
}

func (p *PlanController) ListPlan(c echo.Context) error {
	var params UserIdPath
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	plans, err := p.plan.ListPlan(params.Id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": plans,
	})
}

func (p *PlanController) Create(c echo.Context) error {
	var params model.Plan
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	if err = p.plan.Create(&params); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}

func (p *PlanController) Update(c echo.Context) error {
	var params model.Plan
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	if err = p.plan.Update(&params); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}

func (p *PlanController) Delete(c echo.Context) error {
	var params GetByKey
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	if err = p.plan.Delete(params.Id, params.Key); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": "Success",
	})
}

type GetByKey struct {
	Id  string `param:"id"`
	Key string `param:"key"`
}

func (p *PlanController) GetByKey(c echo.Context) error {
	var params GetByKey
	err := c.Bind(&params)
	if err != nil {
		return c.JSON(400, "bad query params")
	}
	plans, err := p.plan.GetByKey(params.Id, params.Key)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, map[string]interface{}{
		"data": plans,
	})
}
