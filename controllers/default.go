package controllers

import (
	"github.com/astaxie/beego"
)

func init(){
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.Error("test1")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
