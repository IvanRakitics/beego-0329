package controllers

import (
	"Demo0726/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	//"encoding/json"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	pics := models.MapPicturesInfo()
	fmt.Printf("%v\n", pics)
	//c.Data["json"], _ = json.Marshal(pics)
	c.Data["json"] = pics
	c.ServeJSON()
}

func (c *IndexController) Hot() {
	start, _ := c.GetInt("start")
	end, _ := c.GetInt("end")
	Hots := models.MapHotsInfo(start, end)
	fmt.Printf("kkkkkkkkkkkk%v - %v\n", start, end)
	//c.Data["json"], _ = json.Marshal(pics)
	c.Data["json"] = Hots
	c.ServeJSON()
}
