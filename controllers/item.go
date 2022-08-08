package controllers

import (
	"Demo0726/models"
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
)

type ItemsController struct {
	beego.Controller
}

type items_res struct {
	Exit_more bool `json:"exit_more"`
	House_list []models.Items `json:"house_list"`
}

func (c *ItemsController) Get() {
   
	start, _ := c.GetInt("startindex")
	total,_ := c.GetInt("items_length")
	keyword := c.GetString("keyword")
	fmt.Printf(" start %v\n", start)
	fmt.Printf(" keyword %v\n", keyword)
	Items := models.MapItemsInfo(start)
	res := items_res{}
	
	if len(Items) > total {
		res.House_list = Items[0:total]
		res.Exit_more = true
	} else {
		res.House_list = Items
		res.Exit_more = false
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *ItemsController) Details() {
   
	items_id,_ := c.GetInt("items_id")

	fmt.Printf(" items_id %v\n", items_id)
	item_details := models.MapItemsDetailsInfo(items_id)

	c.Data["json"] = item_details
	c.ServeJSON()
}
