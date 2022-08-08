package controllers

import (
	"Demo0726/models"
	//"fmt"
	"encoding/json"
	"fmt"

	// "math/rand"

	beego "github.com/beego/beego/v2/server/web"
)

type AddressController struct {
	beego.Controller
}

type Address struct {
	City             string `json:"city"`
	Consignee        string `json:"consignee"`
	County           string `json:"county"`
	Defaultaddress   int    `json:"defaultaddress"`
	Details_addresss string `json:"details_addresss"`
	Id               int    `json:"id"`
	Phone            string `json:"phone"`
	Province         string `json:"province"`
	Type             string `json:"type"`
	User             int    `json:"user"`
}

func (c *AddressController) Get() {
	user, _ := c.GetInt("user")
	AddressInfo := models.MapAddressInfo(user)
	c.Data["json"] = AddressInfo
	c.ServeJSON()
}

func (c *AddressController) Done() {
	re := c.Ctx.Input.RequestBody
	var address Address
	json.Unmarshal(re, &address)
	fmt.Printf("user %v \n", address)
	addressInfo := models.Address{
		User:             address.User,
		Province:         address.Province,
		City:             address.City,
		County:           address.County,
		Details_addresss: address.Details_addresss,
		Phone:            address.Phone,
		Consignee:        address.Consignee,
		Defaultaddress:   address.Defaultaddress,
	}
	fmt.Printf("addressInfo %v \n", addressInfo)
	var res results
	if address.Type == "insert" {
		id := models.InsertAddress(addressInfo)
		res.Id = id
		res.Code = 300
		res.Message = "SUCCESS"
	} else if address.Type == "update" {
		addressInfo.Id = address.Id
		res.Id = address.Id
		flag := models.UpdateAddress(addressInfo)
		if flag {
			res.Code = 500
			res.Message = "SUCCESS"
		} else {
			res.Code = 300
			res.Message = "FAILED"
		}
	}
	c.Data["json"] = res
	c.ServeJSON()
}
