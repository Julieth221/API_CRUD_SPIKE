package main

import (
	"github.com/astaxie/beego/plugins/cors"
	// _ "github.com/julieth221/API_CRUD_SPIKE/API_CRUD_SENSORES/routers"
	"github.com/udistrital/utils_oas/customerrorv2"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/udistrital/utils_oas/customerrorv2"

	_ "github.com/lib/pq"
)

func main() {
	// Configuración CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "GETALL", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))
	//Configuracion desarrollo
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	beego.Run()
}
