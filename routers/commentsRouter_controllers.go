package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:EstadoCivilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GeneroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:GrupoSanguineoPersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaEstadoCivilController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGeneroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaGrupoEtnicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaPerfilProfesionalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:PersonaTipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:RelacionPersonasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoDiscapacidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/personas_crud/controllers:TipoRelacionPersonasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
