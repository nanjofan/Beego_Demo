package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.1111"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
