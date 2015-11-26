package main

import (
	_ "etnet/routers"
	"github.com/astaxie/beego"
	"etnet/tcp"
	"github.com/astaxie/beego/orm"
)


func initDb() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "db/frame.sqlite3")
	orm.Debug = true
}

func main() {
	go tcp.ServerRun()
	beego.Run()
}

