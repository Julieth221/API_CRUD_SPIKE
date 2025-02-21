// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/YAMITHSALC27/API_CRUD_SPIKE/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),

		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),

		beego.NSNamespace("/tipo_suelo",
			beego.NSInclude(
				&controllers.TipoSueloController{},
			),
		),

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

		beego.NSNamespace("/tipo_arroz",
			beego.NSInclude(
				&controllers.TipoArrozController{},
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

		beego.NSNamespace("/metodo_aplicacion_insumo",
			beego.NSInclude(
				&controllers.MetodoAplicacionInsumoController{},
			),
		),

		beego.NSNamespace("/tipo_informe_cultivo",
			beego.NSInclude(
				&controllers.TipoInformeCultivoController{},
			),
		),

		beego.NSNamespace("/Informe_Cultivo",
			beego.NSInclude(
				&controllers.InformeCultivoController{},
			),
		),

		beego.NSNamespace("/Roles-Usuario",
			beego.NSInclude(
				&controllers.Roles_UsuarioController{},
			),
		),

		beego.NSNamespace("/Roles",
			beego.NSInclude(
				&controllers.RolesController{},
			),
		),

		beego.NSNamespace("/Arrendamiento",
			beego.NSInclude(
				&controllers.ArrendamientoController{},
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

		beego.NSNamespace("/Cultivo_Fase",
			beego.NSInclude(
				&controllers.CultivoFaseController{},
			),
		),

		beego.NSNamespace("/Cultivo_Insumo",
			beego.NSInclude(
				&controllers.CultivoInsumoController{},
			),
		),

		beego.NSNamespace("/Tipo_alertas",
			beego.NSInclude(
				&controllers.TipoAlertasController{},
			),
		),

		beego.NSNamespace("/Alertas",
			beego.NSInclude(
				&controllers.AlertasController{},
			),
		),

		beego.NSNamespace("/Alertas_historial",
			beego.NSInclude(
				&controllers.AlertasHistorialController{},
			),
		),

		beego.NSNamespace("/Configuracion_alertas",
			beego.NSInclude(
				&controllers.ConfiguracionAlertasController{},
			),
		),

		beego.NSNamespace("/Sensor",
			beego.NSInclude(
				&controllers.SensorController{},
			),
		),

		beego.NSNamespace("/Prueba_sensor",
			beego.NSInclude(
				&controllers.PruebaSensorController{},
			),
		),

		beego.NSNamespace("/Historial_sensor",
			beego.NSInclude(
				&controllers.HistorialSensorController{},
			),
		),

		beego.NSNamespace("/Lectura_sensor",
			beego.NSInclude(
				&controllers.LecturaSensorController{},
			),
		),

		beego.NSNamespace("/Tipo_sensor",
			beego.NSInclude(
				&controllers.TipoSensorController{},
			),
		),

		beego.NSNamespace("/Umbral",
			beego.NSInclude(
				&controllers.UmbralController{},
			),
		),

		beego.NSNamespace("/sensor_tipo_sensor",
			beego.NSInclude(
				&controllers.SensorTipoSensorController{},
			),
		),

		beego.NSNamespace("/Geolocalizacion_sensor",
			beego.NSInclude(
				&controllers.GeolocalizacionSensorController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
