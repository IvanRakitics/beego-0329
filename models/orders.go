package models

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Orders struct {
    Id int `orm:"unique" json:"id"`
    Create_time int64 `json:"create_time"`
    Update_time int64 `json:"update_time"`
    Status int `json:"status"`
	User int `json:"user"`
	Address int `json:"address"`
	Amount float32 `json:"amount"`
	Details []OrderDetails `json:"details"  gorm:"foreignkey:Main;association_foreignkey:Id"`
}

type OrderDetails struct {
	Id int `orm:"unique" json:"id"`
	Main int `json:"main"`
	Item int `json:"item"`
	Quantity float32 `json:"quantity"`
	Rate float32 `json:"rate"`
	Delete_flag int `json:"delete_flag"`
	Product []LineProducts `json:"product"  gorm:"foreignkey:Line;association_foreignkey:Id"`
}

type LineProducts struct {
	Id int `orm:"unique" json:"index"`
	Main int `json:"main"`
	Line int `json:"line"`
	Item int `json:"item"`
	ProductId int `json:"productId"`
	ProductName string `json:"productName"`
	Active int `json:"active"`
}

func InsertOrders(data Orders, details []OrderDetails)(int, bool){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Orders{}, &OrderDetails{}, &LineProducts{})
    db.Create(&data)
	var id []int
	db.Raw("SELECT LAST_INSERT_ID() as id").Pluck("id", &id)
	fmt.Printf("id %v", id[0])
	for _, each := range details {
		each.Main = id[0]
		db.Create(&each)
		var line []int
		db.Raw("SELECT LAST_INSERT_ID() as id").Pluck("id", &line)
		fmt.Printf("line %v", line[0])
		for _, v := range each.Product{
			v.Main = id[0]
			v.Line = line[0]
			db.Create(&v)
		}
	}
	return id[0], true
}