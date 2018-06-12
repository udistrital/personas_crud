package models

import (
	"fmt"
	"time"
)

type TrPersona struct {
	Persona           interface{}
	Identificacion    interface{}
	EstadoCivil       interface{}
	Genero            interface{}
	PerfilProfesional interface{}
	GrupoEtnico       interface{}
	TipoDiscapacidad  interface{}
	ContactoEnte      interface{}
	GrupoSanguineo    interface{}
	Rh                interface{}
}

// GetPersonaByIdFull retrieves FUll info from Persona by Id. Returns error if
// Id doesn't exist
func GetPersonaByIdFull(id interface{}) (v *TrPersona, err error) {
	start := time.Now()
	var p TrPersona
	ch := make(chan interface{})
	switch t := id.(type) {
	case int:
		go GetPersonaByIdOnCh(id.(int), ch)
		p.Persona = <-ch
	case string:
		go GetPersonaByUserIdOnCh(id.(string), ch)
		p.Persona = <-ch
	default:
		_ = t
		return
	}
	if p.Persona != nil {
		go GetIdentificacionByIdEnte(p.Persona.(Persona).Ente, ch)
		p.Identificacion = <-ch
		go GetContactoEnteByIdEnte(p.Persona.(Persona).Ente, ch)
		p.ContactoEnte = <-ch
		go GetGeneroPersonaByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.Genero = <-ch
		go GetPersonaEstadoCivilByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.EstadoCivil = <-ch
		go GetPersonaPerfilProfesionalByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.PerfilProfesional = <-ch
		go GetPersonaGrupoEtnicoByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.GrupoEtnico = <-ch
		go GetPersonaTipoDiscapacidadByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.TipoDiscapacidad = <-ch
		go GetPersonaGrupoSanguineoByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.GrupoSanguineo = <-ch
		go GetPersonaRhByIdPersonaOnCh(p.Persona.(Persona).Id, ch)
		p.Rh = <-ch
	} else {
		return
	}

	close(ch)
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	return &p, nil

}
