// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_informe_cultivo",
			beego.NSInclude(
				&controllers.TipoInformeCultivoController{},
			),
		),

		beego.NSNamespace("/Cultivo_Fase",
			beego.NSInclude(
				&controllers.CultivoFaseController{},
			),
		),

		beego.NSNamespace("/Estado_fenologico_cultivo",
			beego.NSInclude(
				&controllers.EstadoFenologicoCultivoController{},
			),
		),

		beego.NSNamespace("/metodo_siembra",
			beego.NSInclude(
				&controllers.MetodoSiembraController{},
			),
		),

		beego.NSNamespace("/metodo_aplicacion_insumo",
			beego.NSInclude(
				&controllers.MetodoAplicacionInsumoController{},
			),
		),

		beego.NSNamespace("/Registro_Cultivo",
			beego.NSInclude(
				&controllers.RegistroCultivoController{},
			),
		),

		beego.NSNamespace("/Detalles_Cultivo",
			beego.NSInclude(
				&controllers.DetallesCultivoController{},
			),
		),

		beego.NSNamespace("/Fase_Cultivo",
			beego.NSInclude(
				&controllers.FaseCultivoController{},
			),
		),

		beego.NSNamespace("/Info_Fase_Cultivo",
			beego.NSInclude(
				&controllers.InfoFaseCultivoController{},
			),
		),

		beego.NSNamespace("/Insumo",
			beego.NSInclude(
				&controllers.InsumoController{},
			),
		),

		beego.NSNamespace("/Informe_Cultivo",
			beego.NSInclude(
				&controllers.InformeCultivoController{},
			),
		),

		beego.NSNamespace("/Cultivo_Insumo",
			beego.NSInclude(
				&controllers.CultivoInsumoController{},
			),
		),

		beego.NSNamespace("/tipo_insumo",
			beego.NSInclude(
				&controllers.TipoInsumoController{},
			),
		),

		beego.NSNamespace("/categoria_insumo",
			beego.NSInclude(
				&controllers.CategoriaInsumoController{},
			),
		),

		beego.NSNamespace("/tipo_arroz",
			beego.NSInclude(
				&controllers.TipoArrozController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
