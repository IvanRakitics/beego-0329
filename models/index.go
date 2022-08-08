package models

import (
	//"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


type Index_pics struct {
    Id int `orm:"unique" json:"id"`
    Img string `json:"img"`
}

func MapPicturesInfo() ([]Index_pics){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Index_pics{})
    var r []Index_pics
    db.Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    return r
}

type Hots struct {
    Id int `orm:"unique" json:"id"`
    Title string `json:"title"`
    Img string `json:"img"`
    Link string `json:"link"`
}

func MapHotsInfo(start int, end int) ([]Hots){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Hots{})
    var r []Hots
    //db.Find(&r) //条件查找所有
    db.Where("Id>=? and Id<=?", start, end).Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    return r
}