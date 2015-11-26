package routers

import (
	"etnet/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/gis", &controllers.GisController{})
	beego.Router("/device", &controllers.DeviceController{})
	beego.Router("/device/params", &controllers.DeviceController{}, "Post:GetParams")

}
