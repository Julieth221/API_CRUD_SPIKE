// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/julieth221/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

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

		beego.NSNamespace("/sensor_geolocalizacion",
			beego.NSInclude(
				&controllers.SensorGeolocalizacionController{},
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
