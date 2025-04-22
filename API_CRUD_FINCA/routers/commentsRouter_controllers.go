package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ArrendamientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:FincaParcelaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:GeolocalizacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:ParcelaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:TipoSueloController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_FINCA/controllers:UserArrendatarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
