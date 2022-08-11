package routers

import (
	"Demo0726/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/pictures", &controllers.IndexController{})
	beego.Router("/hot", &controllers.IndexController{}, "get:Hot")

	beego.Router("/login/check", &controllers.UserController{}, "post:Check")
    beego.Router("/register/email", &controllers.UserController{}, "post:SendEmail")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	beego.Router("/address/list", &controllers.AddressController{})
	beego.Router("/address/update", &controllers.AddressController{},"post:Done")

	beego.Router("/cart/list", &controllers.CartController{})
	beego.Router("/cart/update", &controllers.CartController{}, "post:Update")

	beego.Router("/citys", &controllers.CityController{})

	beego.Router("/items", &controllers.ItemsController{})
	beego.Router("item/details", &controllers.ItemsController{}, "get:Details")

	beego.Router("/category/title", &controllers.CategoryController{})
	beego.Router("/category/list", &controllers.CategoryController{}, "get:Details")

	beego.Router("order/events", &controllers.OrdersController{}, "post:Done")
}
