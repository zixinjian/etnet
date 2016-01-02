package routers

import (
	"etnet/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/gis", &controllers.GisController{})
	beego.Router("/device", &controllers.DeviceController{})
	beego.Router("/device/params", &controllers.DeviceController{}, "Post:GetStatus")
	beego.Router("/device/operate", &controllers.DeviceController{}, "Post:Operate")

}
