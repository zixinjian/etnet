package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["UserName"] = "Test"
	c.TplNames = "main.tpl"
}
