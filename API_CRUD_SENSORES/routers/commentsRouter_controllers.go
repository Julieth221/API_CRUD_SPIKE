package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:AlertasHistorialController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:GeolocalizacionSensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:HistorialSensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:LecturaSensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:PruebaSensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:SensorGeolocalizacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoAlertasController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:TipoSensorController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"] = append(beego.GlobalControllerRouter["github.com/YAMITHSALC27/API_CRUD_SPIKE/API_CRUD_SPIKE/API_CRUD_SENSORES/controllers:UmbralController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
