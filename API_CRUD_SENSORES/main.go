package main

import (
	_ "github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/customerrorv2"

	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	beego.Run()
}
