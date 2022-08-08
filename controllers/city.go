package controllers

import (
	"Demo0726/models"
	beego "github.com/beego/beego/v2/server/web"
)

type CityController struct {
	beego.Controller
}

func (c *CityController) Get() {

	Citys := models.MapCitysInfo()

	c.Data["json"] = Citys
	c.ServeJSON()
}
