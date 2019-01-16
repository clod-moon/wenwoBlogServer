package routers

import (
	"wenwoBlogServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/signup", &controllers.UserController{}, "POST:Signup")
	beego.Router("/login", &controllers.UserController{}, "POST:Login")
	beego.Router("/getarticle", &controllers.UserController{}, "POST:GetArticle")
}
