// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Finca",
			beego.NSInclude(
				&controllers.FincaController{},
			),
		),

		beego.NSNamespace("/Parcela",
			beego.NSInclude(
				&controllers.ParcelaController{},
			),
		),

		beego.NSNamespace("/Arrendamiento",
			beego.NSInclude(
				&controllers.ArrendamientoController{},
			),
		),

		beego.NSNamespace("/tipo_suelo",
			beego.NSInclude(
				&controllers.TipoSueloController{},
			),
		),

		beego.NSNamespace("/Finca_Parcela",
			beego.NSInclude(
				&controllers.FincaParcelaController{},
			),
		),

		beego.NSNamespace("/Geolocalizacion",
			beego.NSInclude(
				&controllers.GeolocalizacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
