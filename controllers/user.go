package controllers

import (
	"Demo0726/models"
	"fmt"
	//"fmt"
	"encoding/json"
	"math/rand"
    //"time"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

type result struct {
	Id int `json:"id"`
	Check bool `json:"check"`
	Token string `json:"token"`
	Nick string `json:"nick"`
}
type User struct {
	Uname string
	Upassword string
	Uphone string
}

func (c *UserController) Check() {
	re := c.Ctx.Input.RequestBody
	var user User
	json.Unmarshal(re, &user)
	fmt.Printf("user %v \n", user)
	Users := models.MapUsersInfo(user.Uphone)
	res := result{}
	res.Check = false
	fmt.Printf("user res %v \n", Users)
	if len(Users) > 0 {
		res.Id = Users[0].Id
        res.Token = GetRandomString(15)
		res.Nick = GetRandomString(15)
		if Users[0].User_password == user.Upassword{
			res.Check = true
		}
        go models.UpdateUserInfo(res.Id, res.Token, res.Nick)
	} else {
		res.Id = 0
		res.Token = GetRandomString(15)
		res.Nick = GetRandomString(15)
	}
	c.Data["json"] = res
	c.ServeJSON()
}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
	   result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
 }