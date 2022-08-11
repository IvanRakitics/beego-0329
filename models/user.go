package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"
)

type Users struct {
    Id int `orm:"unique" json:"id"`
    User_name string `json:"user_name"`
    User_phone string `json:"user_phone"`
    User_password string `json:"user_password"`
    User_email string `json:"user_email"`
	Token string `json:"token"`
    Nick string `json:"nick"`
}

func MapUsersInfo(phone string) ([]Users){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Users{})
    var r []Users
    //db.Find(&r) //条件查找所有
    db.Where("user_phone=?", phone).Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    return r
}

func MapUsersInfoByEmail(email string) ([]Users){
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Users{})
    var r []Users
    //db.Find(&r) //条件查找所有
    db.Where("User_email=?", email).Find(&r) //条件查找所有
    //fmt.Printf("%T\n", poolVolumes)
    return r
}

func UpdateUserInfo(id int,token string, nick string, password string) {

    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Users{})
    var u Users
    db.Model(&u).Where("Id = ?", id).Updates(map[string]interface{}{"Token": token, "Nick": nick, "User_password": password})
}

func RegisterUserInfo(token string, email string, exist bool) {
    
    fmt.Printf("token %v email %v exist %v \n", token, email, exist)
    db, err := gorm.Open("mysql", "root:zhou123456+@(120.48.4.168)/journal?charset=utf8mb4&parseTime=True&loc=Local")
    if err!= nil{
        panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Users{})
    if exist {
        var u Users
        db.Model(&u).Where("User_email=?", email).Updates(map[string]interface{}{"Token": token})
    } else {
        var u Users
        u.User_email = email
        u.Token = token
        db.Create(&u)
    }
}