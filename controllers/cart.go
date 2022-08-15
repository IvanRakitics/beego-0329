package controllers

import (
	"Demo0726/models"
	"encoding/json"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type CartController struct {
	beego.Controller
}

type res_cart struct {
	Code  int           `json:"code"`
	Items []models.Cart `json:"items"`
}

type Cart_data struct {
	Label       []int `json:"label"`
	Count       int   `json:"count"`
	User        int   `json:"user"`
	Item        int   `json:"item"`
	Cancel_flag int   `json:"cancel_flag"`
	Delete_flag int   `json:"delete_flag"`
}

type results struct {
	Id      int    `json:"id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Item    Item `json:"item"`
}
type Item struct {
	Id      int    `json:"id"`
}

func (c *CartController) Get() {
	userid,_ := c.GetInt("userid")
	res := res_cart{}
	res.Code = 500
	res.Items = models.MapCartsInfo(userid)
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *CartController) Update() {
	re := c.Ctx.Input.RequestBody
	var cart_data Cart_data
	err := json.Unmarshal(re, &cart_data)
	fmt.Println(err)

	start_time := time.Now()
	fmt.Printf("start_time: %v second \n", start_time)
	var flag bool
	cart_list := models.SelectCart(cart_data.User, cart_data.Item)
	fmt.Printf("cart_list %v \n", cart_list)
	fmt.Println(len(cart_list) > 0)
	if len(cart_list) > 0 {
		flag = models.UpdateCart(
			cart_list[0].Cart_number,
			cart_data.Cancel_flag,
			cart_data.Count,
			cart_data.Delete_flag,
			cart_data.User,
			cart_data.Item,
			cart_data.Label)
	} else {
		flag = models.InsertCart(
			cart_data.Cancel_flag,
			cart_data.Count,
			cart_data.Delete_flag,
			cart_data.User,
			cart_data.Item,
			cart_data.Label)
	}
	fmt.Printf("during: %v second \n", time.Since(start_time))
	res := results{}
	if flag {
		res.Code = 500
		res.Message = "SUCCESS"
	} else {
		res.Code = 300
		res.Message = "FAILED"
	}
	c.Data["json"] = res
	c.ServeJSON()
}
