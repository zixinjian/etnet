package controllers

import (
	//	"fmt"
	"encoding/base64"
	"github.com/astaxie/beego"
	"etnet/models/userMgr"
	"wb/om"
	"wb/st"
	"wb/cs"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	_, ok := this.GetSession(SessionUser).(om.ValueMap)
	if this.Ctx.Input.Url()== "/" && ok{
		this.Redirect("/main", 302)
		return
	}
	redirectUrl := "/main"
	redirectUrlB64 := this.GetString("redirect")
	if redirectUrlB64 != "" {
		redirectUrlDec, err := base64.URLEncoding.DecodeString(redirectUrlB64)
		if err == nil {
			redirectUrl = string(redirectUrlDec)
		}
	}
	this.Data["redirectUrl"] = redirectUrl
	this.TplNames = "login/login.html"
}
func (this *LoginController) Post() {
	username := this.GetString("login_username")
	password := this.GetString("login_password")

	loginRet := cs.JsonResult{}
	if username == "" || password == "" {
		loginRet.Result = "请输入用户名和密码！"
	}
	params := om.Params{
		"username": username,
		"password": password,
	}
	if code, user := userMgr.GetValidUser(params);code == st.Success {
		this.SetSession(SessionUser, user)
		beego.Info("User:%s login success.", username)
		loginRet.Ret = st.Success
		loginRet.Result = "登录成功！"
	}else{
		beego.Error("User:%s login error code: %s", username, code)
		loginRet.Result = "用户名或密码错误！"
	}
	this.Data["json"] = &loginRet
	this.ServeJson()
}

type LogoutController struct {
	BaseController
}

func (this *LogoutController) Get() {
	this.DestroySession()
	this.Redirect("/", 302)
}
