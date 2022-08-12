package models

import (
	//"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Items struct {
	Id int `orm:"unique" json:"id"`
	Sku int `json:"sku"`
	Title string `json:"title"`
	Housetype string `json:"housetype"`
    Price string `json:"price"`
	Renttype string `json:"renttype"`
	Img string `json:"img"`
	Original_price string `json:"original_price"`
    Item_imgs []ItemImgs `json:"item_imgs" gorm:"foreignkey:Item_id;association_foreignkey:Id"`
    Item_introducts []ItemIntroduces `json:"item_introducts" gorm:"foreignkey:Item_id;association_foreignkey:Id"`
    Product_list []Product `json:"product" gorm:"foreignkey:Goods_id;association_foreignkey:Id"`
}

type ItemImgs struct {
    Item_number int `orm:"unique" json:"item_number"`
	Item_id int `json:"item_id"`
	Item_img string `json:"item_img"`
}

type ItemIntroduces struct {
    Id int `orm:"unique" json:"id"`
	Item_id int `json:"item_id"`
	Item_introduct string `json:"item_introduct"`
}

func MapItemsInfo(start int) ([]Items){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Items{})
    var r []Items
    //db.Find(&r) //条件查找所有
    db.Where("Id>?", start).Find(&r) //条件查找所有
    fmt.Printf("%v\n", r)
    return r
}

func GetItemsInfo(id int) (Items){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Items{})
    var r Items
    //db.Find(&r) //条件查找所有
    db.Where("Id = ?", id).Find(&r) //条件查找所有
    fmt.Printf("%v\n", r)
    return r
}

func MapItemsDetailsInfo(itemId int) (*Items){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Items{}, &ItemImgs{}, &ItemIntroduces{}, &Product{})
    var d = new(Items)
    //db.Find(&r) //条件查找所有
    db.Where("Id=?", itemId).Preload("Item_imgs").Preload("Item_introducts").Preload("Product_list").Find(&d) //条件查找所有
    fmt.Printf("%v\n", d)
    return d
}