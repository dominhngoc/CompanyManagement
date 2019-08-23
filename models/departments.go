package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Departments struct {
	Id       int    `orm:"column(id);pk"`
	Name     string `orm:"column(name);size(255);null"`
	Quantity string `orm:"column(quantity);size(255);null"`
	Employees []*Employees `orm:"reverse(many)"`
}

func (t *Departments) TableName() string {
	return "departments"
}

func init() {
	orm.RegisterModel(new(Departments),new(Employees))
}
func GetRelateDep() {
	o:=orm.NewOrm()
	dep := &Departments{}
	o.QueryTable("Departments").Filter("Id", 1).RelatedSel().One(dep)
	// Get Profile automatically
	fmt.Println("dep.employess")
	fmt.Println(dep.Employees)
}
// AddDepartments insert a new Departments into database and returns
// last inserted Id on success.
func AddDepartments(m *Departments) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDepartmentsById retrieves Departments by Id. Returns error if
// Id doesn't exist
func GetDepartmentsById(id int) (v *Departments, err error) {
	o := orm.NewOrm()
	v = &Departments{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDepartments retrieves all Departments matches certain condition. Returns empty list if
// no records exist
func GetAllDepartments(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Departments))
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

	var l []Departments
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

// UpdateDepartments updates Departments by Id and returns error if
// the record to be updated doesn't exist
func UpdateDepartmentsById(m *Departments) (err error) {
	o := orm.NewOrm()
	v := Departments{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDepartments deletes Departments by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDepartments(id int) (err error) {
	o := orm.NewOrm()
	v := Departments{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Departments{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
