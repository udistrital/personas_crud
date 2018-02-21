// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/personas_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/genero",
			beego.NSInclude(
				&controllers.GeneroController{},
			),
		),

		beego.NSNamespace("/persona_genero",
			beego.NSInclude(
				&controllers.PersonaGeneroController{},
			),
		),

		beego.NSNamespace("/grupo_etnico",
			beego.NSInclude(
				&controllers.GrupoEtnicoController{},
			),
		),

		beego.NSNamespace("/persona_grupo_etnico",
			beego.NSInclude(
				&controllers.PersonaGrupoEtnicoController{},
			),
		),

		beego.NSNamespace("/estado_civil",
			beego.NSInclude(
				&controllers.EstadoCivilController{},
			),
		),

		beego.NSNamespace("/persona_estado_civil",
			beego.NSInclude(
				&controllers.PersonaEstadoCivilController{},
			),
		),

		beego.NSNamespace("/tipo_discapacidad",
			beego.NSInclude(
				&controllers.TipoDiscapacidadController{},
			),
		),

		beego.NSNamespace("/persona_tipo_discapacidad",
			beego.NSInclude(
				&controllers.PersonaTipoDiscapacidadController{},
			),
		),

		beego.NSNamespace("/perfil_profesional",
			beego.NSInclude(
				&controllers.PerfilProfesionalController{},
			),
		),

		beego.NSNamespace("/persona_perfil_profesional",
			beego.NSInclude(
				&controllers.PersonaPerfilProfesionalController{},
			),
		),

		beego.NSNamespace("/tipo_relacion_personas",
			beego.NSInclude(
				&controllers.TipoRelacionPersonasController{},
			),
		),

		beego.NSNamespace("/persona",
			beego.NSInclude(
				&controllers.PersonaController{},
			),
		),

		beego.NSNamespace("/relacion_personas",
			beego.NSInclude(
				&controllers.RelacionPersonasController{},
			),
		),

		beego.NSNamespace("/tipo_ente",
			beego.NSInclude(
				&controllers.TipoEnteController{},
			),
		),

		beego.NSNamespace("/contacto_ente",
			beego.NSInclude(
				&controllers.ContactoEnteController{},
			),
		),

		beego.NSNamespace("/identificacion",
			beego.NSInclude(
				&controllers.IdentificacionController{},
			),
		),

		beego.NSNamespace("/tipo_identificacion",
			beego.NSInclude(
				&controllers.TipoIdentificacionController{},
			),
		),

		beego.NSNamespace("/tipo_contacto",
			beego.NSInclude(
				&controllers.TipoContactoController{},
			),
		),

		beego.NSNamespace("/ente",
			beego.NSInclude(
				&controllers.EnteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
