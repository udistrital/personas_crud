package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PersonaPerfilProfesional struct {
	Id                int                `orm:"column(id);pk;auto"`
	PerfilProfesional *PerfilProfesional `orm:"column(perfil_profesional);rel(fk)"`
	Persona           *Persona           `orm:"column(persona);rel(fk)"`
}

func (t *PersonaPerfilProfesional) TableName() string {
	return "persona_perfil_profesional"
}

func init() {
	orm.RegisterModel(new(PersonaPerfilProfesional))
}

// AddPersonaPerfilProfesional insert a new PersonaPerfilProfesional into database and returns
// last inserted Id on success.
func AddPersonaPerfilProfesional(m *PersonaPerfilProfesional) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPersonaPerfilProfesionalById retrieves PersonaPerfilProfesional by Id. Returns error if
// Id doesn't exist
func GetPersonaPerfilProfesionalById(id int) (v *PersonaPerfilProfesional, err error) {
	o := orm.NewOrm()
	v = &PersonaPerfilProfesional{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPersonaPerfilProfesional retrieves all PersonaPerfilProfesional matches certain condition. Returns empty list if
// no records exist
func GetAllPersonaPerfilProfesional(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PersonaPerfilProfesional))
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

	var l []PersonaPerfilProfesional
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

// UpdatePersonaPerfilProfesional updates PersonaPerfilProfesional by Id and returns error if
// the record to be updated doesn't exist
func UpdatePersonaPerfilProfesionalById(m *PersonaPerfilProfesional) (err error) {
	o := orm.NewOrm()
	v := PersonaPerfilProfesional{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePersonaPerfilProfesional deletes PersonaPerfilProfesional by Id and returns error if
// the record to be deleted doesn't exist
func DeletePersonaPerfilProfesional(id int) (err error) {
	o := orm.NewOrm()
	v := PersonaPerfilProfesional{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PersonaPerfilProfesional{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetPersonaPerfilProfesionalByIdPersonaOnCh retrieves PerfilProfesional by PersonaId. Returns error if
// Id doesn't exist
func GetPersonaPerfilProfesionalByIdPersonaOnCh(id int, c chan interface{}) (err error) {
	o := orm.NewOrm()
	var pg []PersonaPerfilProfesional
	qs := o.QueryTable(new(PersonaPerfilProfesional)).RelatedSel("perfil_profesional")
	qs.Filter("persona", id).All(&pg, "perfil_profesional")
	var perfiles []PerfilProfesional
	for _, vp := range pg {
		perfiles = append(perfiles, *vp.PerfilProfesional)
	}
	c <- perfiles
	return nil
}
