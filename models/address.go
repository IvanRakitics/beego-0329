package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Address struct {
	Id               int    `orm:"unique" json:"id"`
	User             int    `json:"user"`
	Province         string `json:"province"`
	City             string `json:"city"`
	County           string `json:"county"`
	Details_addresss string `json:"details_addresss"`
	Phone            string `json:"phone"`
	Consignee        string `json:"consignee"`
	Defaultaddress   int    `json:"defaultaddress"`
}

func MapAddressInfo(user int) []Address {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Address{})
	var r []Address
	//db.Find(&r) //条件查找所有
	db.Where("User = ?", user).Find(&r) //条件查找所有
	//fmt.Printf("%T\n", poolVolumes)
	return r
}

func GetDefaultAddressInfo(user int) (Address) {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Address{})
	var r []Address
	db.Where("User = ?", user).Find(&r) //条件查找所有
	if len(r) == 0 {
		return Address{}
	}
	var address Address
	exist := false
	for _,each := range r {
		if each.Defaultaddress == 1 {
			exist = true
			address = each
		}
	}
	if !exist {
		return r[0]
	}
	return address
}

func InsertAddress(addressInfo Address) int {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var address Address
	db.AutoMigrate(&Address{})
	if addressInfo.Defaultaddress == 1 {
		db.Model(&address).Where("User = ?", addressInfo.User).Updates(map[string]interface{}{"defaultaddress": 0})
	}
	db.Create(&addressInfo)
	var id []int
	db.Raw("SELECT LAST_INSERT_ID() as id").Pluck("id", &id)
	fmt.Printf("id %v", id[0])
	return id[0]
}

func UpdateAddress(addressInfo Address) bool {
	db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var address Address
	db.AutoMigrate(&Address{})
	if addressInfo.Defaultaddress == 1 {
        fmt.Printf("User %v \n", addressInfo.User)
		err1 := db.Model(&address).Where("User = ?", addressInfo.User).Updates(map[string]interface{}{"defaultaddress": 0})
		if err1.Error != nil {
			fmt.Printf("err1 %v", err1.Error)
			return false
		}
	}
	err2 := db.Model(&address).Where("Id = ?", addressInfo.Id).Updates(
		map[string]interface{}{
			"Province":         addressInfo.Province,
			"City":             addressInfo.City,
			"County":           addressInfo.County,
			"Details_addresss": addressInfo.Details_addresss,
			"Phone":            addressInfo.Phone,
			"Consignee":        addressInfo.Consignee,
			"Defaultaddress":   addressInfo.Defaultaddress,
		})
	if err2.Error != nil {
		fmt.Printf("err2 %v", err2.Error)
		return false
	}
	return true
}
