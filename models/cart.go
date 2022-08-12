package models

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Cart struct {
	Cart_number  int           `json:"cart_number" gorm:"-;primary_key;AUTO_INCREMENT"`
	User_id      int           `json:"user_id"`
	Sku          string        `json:"sku"`
	Goods_id     string        `json:"goods_id"`
	Product_id   string        `json:"product_id"`
	Count        int           `json:"count"`
	Delete_flag  int           `json:"delete_flag"`
	Cancel_flag  int           `json:"cancel_flag"`
	Img          string        `json:"img"`
	Product      []CartProduct `json:"product_name" gorm:"foreignkey:Cart_number;association_foreignkey:Cart_number"`
	Item_details Items         `json:"item_details" gorm:"foreignkey:Goods_id;association_foreignkey:Id"`
}

type CartProduct struct {
	Index       int     `orm:"unique" json:"index"`
	Cart_number int     `json:"cart_number"`
	User_id     int     `json:"user_id"`
	Goods_id    int     `json:"goods_id"`
	Product_id  int     `json:"product_id"`
	Active_flag int     `json:"active_flag"`
	Product     Product `json:"product_name" gorm:"foreignkey:Number;association_foreignkey:Product_id"`
}

type Product struct {
	Number       int    `json:"number"`
	Goods_id     int    `json:"goods_id"`
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Title        string `json:"title"`
	Cart_img     string `json:"cart_img"`
}

func MapCartsInfo(User int) []Cart {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Cart{}, &Product{}, &Items{}, &CartProduct{})
	var r []Cart
	//db.Find(&r) //条件查找所有
	//db.Find(&r) //条件查找所有
	db.Debug().Where("User_id = ?", User).Preload("Product.Product").Preload("Product").Preload("Item_details").Find(&r)
	//fmt.Printf("%T\n", poolVolumes)
	return r
}

func SelectCart(user int, item int) []Cart {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Cart{})
	var cart_list []Cart
	db.Where("User_id=? and Goods_id=?", user, item).Find(&cart_list)
	return cart_list
}

func UpdateCart(cart_number int, cancel_flag int, count int, delete_flag int, user int, item int, label []int) bool {
	fmt.Printf("UpdateCart %v \n", cart_number)
	fmt.Printf("UpdateCart delete_flag %v \n", delete_flag)
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Cart{}, &Product{})
	var cart Cart
	var cartProduct CartProduct
	err1 := db.Model(&cart).Where("Cart_number = ?", cart_number).Updates(map[string]interface{}{"cancel_flag": cancel_flag, "count": count, "delete_flag": delete_flag})
	//var cartProductList []CartProduct
	if err1.Error != nil {
		fmt.Printf("err1 %v", err1.Error)
		return false
	}
	err2 := db.Model(&cartProduct).Where("cart_number = ?", cart_number).Delete(&cartProduct)
	if err2.Error != nil {
		fmt.Printf("err2 %v", err2.Error)
		return false
	}
	for _, each := range label {
		fmt.Println(each)
		var cartProduct = new(CartProduct)
		cartProduct.Cart_number = cart_number
		cartProduct.User_id = user
		cartProduct.Goods_id = item
		cartProduct.Product_id = each
		cartProduct.Active_flag = 1
		db.Create(&cartProduct)
	}

	return true
}

func InsertCart(cancel_flag int, count int, delete_flag int, user int, item int, label []int) bool {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	cart := Cart{User_id: user, Count: count, Cancel_flag: cancel_flag, Goods_id: strconv.Itoa(item)}
	db.AutoMigrate(&Cart{}, &Items{}, &Product{})
	items := MapItemsDetailsInfo(item)
	var product Product
	db.Where("Number=?", label[0]).Find(&product)
	cart.Sku = strconv.Itoa(items.Sku)
	cart.Img = product.Cart_img
	fmt.Println(cart)
	err2 := db.Create(&cart)

	var id []int
	db.Raw("SELECT LAST_INSERT_ID() as id").Pluck("id", &id)
	fmt.Printf("id %v", id[0])
	if err2.Error != nil {
		fmt.Printf("err2 %v", err2.Error)
		fmt.Printf("cart %v", cart)
		return false
	} else {
		for _, each := range label {
			fmt.Println(each)
			var cartProduct = new(CartProduct)
			cartProduct.Cart_number = id[0]
			cartProduct.User_id = user
			cartProduct.Goods_id = item
			cartProduct.Product_id = each
			cartProduct.Active_flag = 1
			db.Create(&cartProduct)
		}
	}
	return true
}

// func MapProductsInfos(index int) (Product) {
// 	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	db.AutoMigrate(&Product{})
// 	var r Product
//     //db.Find(&r) //条件查找所有
//     db.Where("Number = ?", index).Find(&r) //条件查找所有
//     fmt.Printf("%v\n", r)
//     return r
// }