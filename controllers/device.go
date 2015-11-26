package controllers
import (
	"github.com/astaxie/beego"
	"etnet/models/device"
)


type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Get() {
	sn := c.GetString("sn")
	if sn==""{
		c.Ctx.WriteString("无此ID")
		return
	}
	c.Data["DeviceName"] = "发电机组A"
	c.Data["DeviceId"] = sn
}


func (c *DeviceController)GetParams(){
	sn := c.GetString("sn")
	if sn==""{
		c.Ctx.WriteString("无此ID")
		return
	}
	c.Data["json"] = device.GetParams(sn)
	c.ServeJson()
}