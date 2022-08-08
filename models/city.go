package models

import (
	//"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type City struct {
    Id int `orm:"unique" json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

func MapCitysInfo() ([]City){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&City{})
    var r []City
    //db.Find(&r) //条件查找所有
    db.Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    return r
}