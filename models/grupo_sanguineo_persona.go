package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type GrupoSanguineoPersona struct {
	Id                int      `orm:"column(id);pk;auto"`
	FactorRh          string   `orm:"column(factor_rh)"`
	GrupoSanguineo    string   `orm:"column(grupo_sanguineo)"`
	Persona           *Persona `orm:"column(persona);rel(fk)"`
	FechaCreacion     string   `orm:"column(fecha_creacion);null"`
	FechaModificacion string   `orm:"column(fecha_modificacion);null"`
}

func (t *GrupoSanguineoPersona) TableName() string {
	return "grupo_sanguineo_persona"
}

func init() {
	orm.RegisterModel(new(GrupoSanguineoPersona))
}

// AddGrupoSanguineoPersona insert a new GrupoSanguineoPersona into database and returns
// last inserted Id on success.
func AddGrupoSanguineoPersona(m *GrupoSanguineoPersona) (id int64, err error) {
	m.FechaCreacion = time_bogota.TiempoBogotaFormato()
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGrupoSanguineoPersonaById retrieves GrupoSanguineoPersona by Id. Returns error if
// Id doesn't exist
func GetGrupoSanguineoPersonaById(id int) (v *GrupoSanguineoPersona, err error) {
	o := orm.NewOrm()
	v = &GrupoSanguineoPersona{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGrupoSanguineoPersona retrieves all GrupoSanguineoPersona matches certain condition. Returns empty list if
// no records exist
func GetAllGrupoSanguineoPersona(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GrupoSanguineoPersona))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []GrupoSanguineoPersona
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateGrupoSanguineoPersona updates GrupoSanguineoPersona by Id and returns error if
// the record to be updated doesn't exist
func UpdateGrupoSanguineoPersonaById(m *GrupoSanguineoPersona) (err error) {
	o := orm.NewOrm()
	v := GrupoSanguineoPersona{Id: m.Id}
	m.FechaModificacion = time_bogota.TiempoBogotaFormato()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, "FactorRh", "GrupoSanguineo", "Persona", "FechaModificacion"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGrupoSanguineoPersona deletes GrupoSanguineoPersona by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGrupoSanguineoPersona(id int) (err error) {
	o := orm.NewOrm()
	v := GrupoSanguineoPersona{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GrupoSanguineoPersona{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetPersonaGrupoSanguineoByIdPersonaOnCh retrieves EstadoCivil by PersonaId. Returns error if
// Id doesn't exist
func GetPersonaGrupoSanguineoByIdPersonaOnCh(id int, c chan interface{}) (err error) {

	o := orm.NewOrm()
	var pg GrupoSanguineoPersona
	qs := o.QueryTable(new(GrupoSanguineoPersona))
	fmt.Println(qs)
	qs.Filter("persona", id).All(&pg, "GrupoSanguineo", "FactorRh")
	c <- pg.GrupoSanguineo

	return nil
}

// GetPersonaGrupoSanguineoByIdPersonaOnCh retrieves EstadoCivil by PersonaId. Returns error if
// Id doesn't exist
func GetPersonaRhByIdPersonaOnCh(id int, c chan interface{}) (err error) {

	o := orm.NewOrm()
	var pg GrupoSanguineoPersona
	qs := o.QueryTable(new(GrupoSanguineoPersona))

	qs.Filter("persona", id).All(&pg, "FactorRh")
	c <- pg.FactorRh

	return nil
}
