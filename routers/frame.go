package routers
import (
"github.com/astaxie/beego"
"etnet/controllers"
)


func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
}