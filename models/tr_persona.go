package models

type TrPersona struct {
	Persona           interface{}
	EstadoCivil       interface{}
	Genero            interface{}
	PerfilProfesional interface{}
	GrupoEtnico       interface{}
	TipoDiscapacidad  interface{}
}

// GetPersonaByIdFull retrieves FUll info from Persona by Id. Returns error if
// Id doesn't exist
func GetPersonaByIdFull(id int) (v *TrPersona, err error) {

	var p TrPersona
	ch := make(chan interface{}, 7)
	go GetPersonaByIdOnCh(id, ch)
	p.Persona = <-ch
	// fmt.Println(id)
	// //fmt.Println("av", <-ch)
	// fmt.Println("v", v)
	go GetGeneroPersonaByIdPersonaOnCh(id, ch)
	p.Genero = <-ch
	go GetPersonaEstadoCivilByIdPersonaOnCh(id, ch)
	p.EstadoCivil = <-ch
	go GetPersonaPerfilProfesionalByIdPersonaOnCh(id, ch)
	p.PerfilProfesional = <-ch
	go GetPersonaGrupoEtnicoByIdPersonaOnCh(id, ch)
	p.GrupoEtnico = <-ch
	go GetPersonaTipoDiscapacidadByIdPersonaOnCh(id, ch)
	p.TipoDiscapacidad = <-ch
	print("antes", v)
	close(ch)
	return &p, nil

}
