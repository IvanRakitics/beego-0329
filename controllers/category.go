package controllers

import (
	"Demo0726/models"
	//"fmt"
	beego "github.com/beego/beego/v2/server/web"
	//"encoding/json"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	res := models.MapCategorysInfo()
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *CategoryController) Details() {
	res := models.MapCategoryDetailsInfo()
	c.Data["json"] = res
	c.ServeJSON()
}