package controllers
import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *DeviceController) SendJson(jsonObject interface{}) {
	c.Data["json"] = &jsonObject
	c.ServeJson()
}

