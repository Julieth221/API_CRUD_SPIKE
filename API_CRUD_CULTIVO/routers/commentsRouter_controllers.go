package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CategoriaInsumoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoFaseController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:CultivoInsumoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:DetallesCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:EstadoFenologicoCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:FaseCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InfoFaseCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InformeCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:InsumoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoAplicacionInsumoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:MetodoSiembraController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:RegistroCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoArrozController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInformeCultivoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_CULTIVO/controllers:TipoInsumoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
