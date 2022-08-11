package controllers

import (
	"Demo0726/models"
	"encoding/json"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type OrdersController struct {
	beego.Controller
}

type DataInfo struct {
	Amount  float32        `json:"amount"`
	Status  int            `json:"status"`
	User    int            `json:"user"`
	Details []OrderDetails `json:"details"`
}

type OrderDetails struct {
	Item     int     `json:"item"`
	Product  []int   `json:"product"`
	Quantity float32 `json:"quantity"`
	Rate     float32 `json:"rate"`
}

type OrdersType struct {
	Eventstype string `json:"type" `
}

type OrdersInfo struct {
	Events OrdersType `json:"events" `
	Data   DataInfo   `json:"data" `
}

func (c *OrdersController) Done() {
	res := results{Code: 300}
    res.Item = Item{Id: 0}

	re := c.Ctx.Input.RequestBody
	var ordersInfo OrdersInfo
	err := json.Unmarshal(re, &ordersInfo)

	var orders models.Orders
	orders.Amount = float32(ordersInfo.Data.Amount)
	orders.Status = ordersInfo.Data.Status
	orders.User = ordersInfo.Data.User
	defaultAddress := models.GetDefaultAddressInfo(orders.User)
	orders.Address = defaultAddress.Id
	orders.Create_time = time.Now().Unix()
	orders.Update_time = time.Now().Unix()

	details := make([]models.OrderDetails, len(ordersInfo.Data.Details))
	for _, each := range ordersInfo.Data.Details {
		line := new(models.OrderDetails)
		line.Delete_flag = 0
		line.Item = each.Item
		line.Quantity = each.Quantity
		line.Rate = each.Rate
		products := make([]models.LineProducts, len(each.Product))
		for index, v := range each.Product {
			var lineproducts models.LineProducts
			lineproducts.ProductId = v
			lineproducts.Active = 1
			lineproducts.Item = each.Item
			products[index] = lineproducts
		}
		line.Product = products
		details = append(details, *line)
	}
	fmt.Printf("orders %v details %v", orders, details)

	if ordersInfo.Events.Eventstype == "create" {
		res.Item.Id,_ = models.InsertOrders(orders, details)
		res.Code = 500
	}
	fmt.Println(err)
	fmt.Println(ordersInfo)
	c.Data["json"] = res
	c.ServeJSON()
}
