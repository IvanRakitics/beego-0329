package models

import (
	"fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "sort"
)

type CategoryList []Categorys

func (c CategoryList) Len() int {    // 重写 Len() 方法
    return len(c)
}

func (c CategoryList) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
    return c[j].Category_id < c[i].Category_id 
}

func (a CategoryList) Swap(i, j int){     // 重写 Swap() 方法
    a[i], a[j] = a[j], a[i]
}

type Categorys struct {
    Block_id int `orm:"unique" json:"block_id"`
	Category_id int `json:"category_id"`
    Category_name string `json:"category_name"`
	Is_default int `json:"is_default"`
    Log_code string `json:"log_code"`
    Details []CategoryLists `json:"details" gorm:"foreignkey:Category_id;association_foreignkey:Category_id"`
}

type CategoryLists struct {
    Id int `orm:"unique" json:"id"`
    Category_id int `json:"category_id"`
    Is_show string `json:"is_show"`
    Is_expand int `json:"is_expand"`
    View_type string `json:"view_type"`
    Items []CategoryDetails `json:"items" gorm:"foreignkey:Category_list;association_foreignkey:Id"`
}

type CategoryDetails struct {
    Id int `orm:"unique" json:"id"`
    Category_list int `json:"category_list"`
    Category_title string `json:"category_title"`
    Item_id int `json:"item_id"`
    Img string `json:"img"`
    Product Items `json:"product" gorm:"foreignkey:Id;association_foreignkey:Item_id"`
}

func MapCategorysInfo() ([]Categorys){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Categorys{})
    var r []Categorys
    //db.Find(&r) //条件查找所有
    db.Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    sort.Sort(sort.Reverse(CategoryList(r))) 
    fmt.Printf("%v\n", r)
    return r
}

func MapCategoryDetailsInfo() ([]Categorys){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Categorys{}, &CategoryLists{}, &CategoryDetails{}, &Items{})
    var r []Categorys
    //db.Find(&r) //条件查找所有
    db.Preload("Details").Preload("Details.Items").Preload("Details.Items.Product").Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    sort.Sort(sort.Reverse(CategoryList(r))) 
    fmt.Printf("%v\n", r)
    return r
}

