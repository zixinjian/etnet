package controllers
import (
	"github.com/astaxie/beego"
	"wb/st"
	"etnet/models/statusMgr"
	"wb/ut"
	"wb/cc"
)


type DeviceController struct {
	beego.Controller
}

func (c *DeviceController) Get() {
	sn := c.GetString(cc.Sn)
	if sn==""{
		beego.Error("Get ", st.ParamSnIsNone)
		c.Ctx.WriteString(st.ParamSnIsNone)
		return
	}
	c.Data["DeviceName"] = "发电机组A"
	c.Data["DeviceId"] = sn
}


func (c *DeviceController)GetStatus(){
	sn := c.GetString(cc.Sn)
	if sn==""{
		beego.Error("GetStatus ", st.ParamSnIsNone)
		c.Ctx.WriteString("{}")
		return
	}
	iSn, err := ut.StrTo(sn).Int64()
	if err != nil{
		beego.Error("GetStatus ", st.SnError)
		c.Ctx.WriteString("{}")
		return
	}
	c.Data["json"] = statusMgr.GetStatus(iSn)
	c.ServeJson()
}