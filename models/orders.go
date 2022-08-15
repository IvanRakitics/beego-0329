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
	AddressInfo Address `json:"addressInfo" gorm:"foreignkey:Id;association_foreignkey:Address"`
	Amount float32 `json:"amount"`
	Details []OrderDetails `json:"details" gorm:"foreignkey:Main;association_foreignkey:Id"`
}

type OrderDetails struct {
	Id int `orm:"unique" json:"id"`
	Main int `json:"main"`
	Item int `json:"item"`
	Item_title string `json:"item_title"`
	Cart_img string `json:"cart_img"`
	Quantity float32 `json:"quantity"`
	Rate float32 `json:"rate"`
	Delete_flag int `json:"delete_flag"`
	Product []LineProducts `json:"product"  gorm:"foreignkey:Line;association_foreignkey:Id"`
}

type LineProducts struct {
	Id int `gorm:"-;primary_key;AUTO_INCREMENT" json:"index"`
	Order int `json:"order"`
	Line int `json:"line"`
	Item int `json:"item"`
	ProductId int `json:"productId"`
	Product Product `json:"product"  gorm:"foreignkey:Number;association_foreignkey:ProductId"`
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
	fmt.Printf("id %v %T \n", id[0], id[0])
	var MainId int = id[0]
	for _, each := range details {
		each.Main = MainId
		//db.Create(&each)
		//var line []int
		//db.Raw("SELECT LAST_INSERT_ID() as id").Pluck("id", &line)
		//fmt.Printf("line %v Main %v \n", line[0], each.Main)
		for i, v := range each.Product{
			v.Order = MainId
			//v.Line = line[0]
			//fmt.Printf("LineProducts %v %v %v\n", v, v.Order, MainId)
			//db.Create(&v)
			each.Product[i] = v
		}
		db.Create(&each)
	}
	return MainId, true
}

func UpdatedOrders(data Orders, details []OrderDetails)(int, bool){
    
	fmt.Printf("UpdatedOrders %v %v id: %v\n", data, details, data.Id)

	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Orders{}, &OrderDetails{}, &LineProducts{})

	db.Model(&Orders{}).Where("Id = ?", data.Id).Updates(map[string]interface{}{
		"Status": data.Status,
		"Address": data.Address,
		"Amount": data.Amount,
		"Update_time": data.Update_time,
	})

	var orderDetails OrderDetails
	err1 := db.Model(&orderDetails).Where("Main = ?", data.Id).Delete(&orderDetails)

	if err1.Error != nil {
		fmt.Printf("err1 %v\n", err1)
		return data.Id,false
	}

	var lineProducts LineProducts
	err2 := db.Model(&lineProducts).Where("`order` = ?", data.Id).Delete(&lineProducts)

	if err2.Error != nil {
		return data.Id,false
	}

	for _, each := range details {
		each.Main = data.Id
		for i, v := range each.Product{
			v.Order = data.Id
			each.Product[i] = v
		}
		db.Create(&each)
	}

	return data.Id,true
}

func MapOrderInfos( id int) (Orders){
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Orders{}, &OrderDetails{}, &LineProducts{}, &Address{}, &Product{})
	var r = new(Orders)
    //db.Find(&r) //条件查找所有
    db.Where("Id=?", id).Preload("AddressInfo").Preload("Details").Preload("Details.Product").Find(&r) //条件查找所有
    fmt.Printf("%v\n", r)
	return *r
}

func MapUserOrdersInfos(user int) ([]Orders){
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Orders{}, &OrderDetails{}, &LineProducts{}, &Address{}, &Product{})
	var r []Orders
	db.Where("User=?", user).Preload("AddressInfo").Preload("Details").Preload("Details.Product").Find(&r) //条件查找所有
    fmt.Printf("%v\n", r)
	return r
}

func AppendOrderInfos( id int){
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&OrderDetails{}, &Product{}, &LineProducts{})

	carts := make(map[int]string)

	var pros []LineProducts
	db.Debug().Where("`order`=?", id).Preload("Product").Find(&pros)

	if len(pros) > 0 {
		for _, each := range pros {
			db.Model(&LineProducts{}).Where("Id = ?", each.Id).Updates(map[string]interface{}{"ProductName": each.Product.Title})
		    carts[each.Line] = each.Product.Cart_img
		}
	}
	fmt.Printf("pros %v \n", pros)
	fmt.Printf("carts %v \n", carts)

	var lines []OrderDetails
	db.Where("Main=?", id).Find(&lines)
	if len(lines) > 0 {
		for _, each := range lines {
			if each.Item > 0 {
				Items := GetItemsInfo(each.Item)
				db.Model(&OrderDetails{}).Where("Id = ?", each.Id).Updates(map[string]interface{}{"Item_title": Items.Title, "Cart_img":carts[each.Id]})
			}
		}
	}
}


