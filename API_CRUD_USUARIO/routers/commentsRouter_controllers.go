package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:Roles_UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/julieth221/API_CRUD_SPIKE/API_CRUD_USUARIO/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
