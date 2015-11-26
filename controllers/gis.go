package controllers
import "github.com/astaxie/beego"


type GisController struct {
	beego.Controller
}

func (c *GisController) Get() {
	c.Data["UserName"] = "Test"
}