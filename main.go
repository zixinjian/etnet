package main

import (
	_ "etnet/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"etnet/tcp/modbus"
	"etnet/tcp/bds"
)


func initDb() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "data/db/main.db")
	orm.Debug = true
}

func main() {
	go modbus.ServerRun()
	go bds.ServerRun()
	initDb()
	beego.Run()
}

