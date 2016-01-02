package main

import (
	_ "etnet/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"etnet/tcp/modbus"
	"etnet/tcp/bds"
	"github.com/astaxie/beego/context"
	"etnet/controllers"
	"encoding/base64"
	"wb/om"
)


func initDb() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "data/db/main.db")
	orm.Debug = true
}
var FilterUser = func(ctx *context.Context) {
	if ctx.Request.RequestURI == "/logout" {
		return
	}
	if ctx.Input.Url() == "/login" {
		return
	}
	_, ok := ctx.Input.Session(controllers.SessionUser).(om.ValueMap)
	beego.Debug("FilterUser need login ok: ", ok, ctx.Input.Session(controllers.SessionUser))
	if !ok && ctx.Request.RequestURI != "/login" {
		beego.Debug("FilterUser need login: ", ctx.Input.Url(), ctx.Input.Uri())
		redirect := ctx.Input.Url()
		redirectB64 := base64.URLEncoding.EncodeToString([]byte(redirect))
		ctx.Redirect(302, "/login?redirect="+redirectB64)
	}
}
func main() {
	go modbus.ServerRun()
	go bds.ServerRun()
	initDb()
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
}

