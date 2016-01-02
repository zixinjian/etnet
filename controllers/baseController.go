package controllers
import (
	"github.com/astaxie/beego"
	"etnet/models/userMgr"
)

type BaseController struct {
	beego.Controller
}

func (c *DeviceController) SendJson(jsonObject interface{}) {
	c.Data["json"] = &jsonObject
	c.ServeJson()
}

func (c *DeviceController) GetSessionUser()userMgr.User{
	return c.GetSession(SessionUser).(userMgr.User)
}

