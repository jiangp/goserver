package main

import (
	_ "loginzcsvr/routers"
	"github.com/astaxie/beego"
	//_ "beeapi/docs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

