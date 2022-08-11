package controllers

import (
	"Demo0726/models"
	"fmt"
	//"fmt"
	"encoding/json"
	"math/rand"
    "time"
	beego "github.com/beego/beego/v2/server/web"
	"gopkg.in/gomail.v2"
	"crypto/tls"
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
	Uemail string
	Token string
	UAccountType string
}

type EmailInfo struct {
	Email string
}

func (c *UserController) Check() {
	re := c.Ctx.Input.RequestBody
	var user User
	json.Unmarshal(re, &user)
	fmt.Printf("user %v UAccountType %v\n", user, user.UAccountType)

	fmt.Printf("flag %v %v\n", user.UAccountType == "email", user.UAccountType == "phone")

	var Users []models.Users
	if user.UAccountType == "email" {
		Users = models.MapUsersInfoByEmail(user.Uname)
	} else if user.UAccountType == "phone" {
		Users = models.MapUsersInfo(user.Uname)
	}
	
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
        go models.UpdateUserInfo(res.Id, res.Token, res.Nick, user.Upassword)
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
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
	   result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
 }

 func (c *UserController) SendEmail() {

	re := c.Ctx.Input.RequestBody
	var emailInfo EmailInfo
	json.Unmarshal(re, &emailInfo)
	fmt.Printf("email %v \n", emailInfo.Email)
	fmt.Printf("emailInfo %v \n", emailInfo)

    res := result{}  

	Users := models.MapUsersInfoByEmail(emailInfo.Email)
    exist := len(Users) > 0
	fmt.Printf("Users %v exist %v\n", Users, exist)

	var Token string = GetRandomString(15)
	go sendTokenToEmail(emailInfo.Email, Token)
	go models.RegisterUserInfo(Token, emailInfo.Email, exist)

	if exist {
		res.Id = Users[0].Id
	}
	res.Token = Token

	c.Data["json"] = res
	c.ServeJSON()
 }

 func sendTokenToEmail(userEmail string, Token string){
	message := `
    <p> Hello,</p>
	
		<p style="text-indent:2em">验证码： %s</p> 
		<p style="text-indent:2em">Best Wishes!</p>
	`
 
	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	//SPZTSVHYCJMPRFXN 163授权码
	host := "smtp.163.com"
	port := 25
	userName := "ivanzhou2021@163.com"
	password := "SPZTSVHYCJMPRFXN"
 
	m := gomail.NewMessage()
	m.SetHeader("From", userName) // 发件人
	// m.SetHeader("From", "alias"+"<"+userName+">") // 增加发件人别名
 
	m.SetHeader("To", userEmail) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	//m.SetHeader("Cc", "******@qq.com")                  // 抄送，可以多个
	//m.SetHeader("Bcc", "******@qq.com")                 // 暗送，可以多个
	m.SetHeader("Subject", "Hello!")                    // 邮件主题
 
	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", fmt.Sprintf(message, Token))
 
	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片
 
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
 
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
 }

 func (c *UserController) Register(){
	re := c.Ctx.Input.RequestBody
	var user User
	json.Unmarshal(re, &user)
	fmt.Printf("user %v \n", user)

	Users := models.MapUsersInfoByEmail(user.Uemail)
    exist := len(Users) > 0
	fmt.Printf("Users %v exist %v\n", Users, exist)

	res := result{}
	res.Check = false
	if exist {
		fmt.Printf("token_input %v token %v flag %v\n", user.Token, Users[0].Token, Users[0].Token == user.Token)
		if Users[0].Token == user.Token {
			res.Id = Users[0].Id
			res.Check = true
			res.Token = user.Token
			res.Nick = GetRandomString(15)
			go models.UpdateUserInfo(res.Id, user.Token, res.Nick, user.Upassword)
		} else {
			res.Id = Users[0].Id
			res.Check = false
		}
		
	} else {
		res.Id = 0
		res.Check = false
	}
	c.Data["json"] = res
	c.ServeJSON()
 }