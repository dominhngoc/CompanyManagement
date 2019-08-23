package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type Employees struct {
	Email string `orm:"column(email);size(255);null"`
	Age   int    `orm:"column(age);null"`
	Name  string `orm:"column(name);size(255);null"`
	Id    int    `orm:"column(id);pk"`
	//Id_dep *int  `orm:"rel(fk)"`
	IdDepartment *Departments `orm:"column(id_department);rel(fk)"`
}

func (t *Employees) TableName() string {
	return "employees"
}



// AddEmployees insert a new Employees into database and returns
// last inserted Id on success.
func AddEmployees(m *Employees) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func AddOneEmployees(m *Employees) {
	fmt.Println("'add one runned")
	o := orm.NewOrm()
	id, err := o.Insert(m)
	_=id
	if err!=nil{
		fmt.Println("success")
	}
	return
}
func AddManyEmployees(m *[]Employees) {
	fmt.Println("'add many runned")
	o := orm.NewOrm()
	successNums, err := o.InsertMulti(100, m)
	_=successNums
	_=err
	return
}
// GetEmployeesById retrieves Employees by Id. Returns error if
// Id doesn't exist
func GetEmployeesById(id int) (v *Employees, err error) {
	o := orm.NewOrm()
	v = &Employees{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetRelateModel(){
	var dep Departments
	o:=orm.NewOrm()
	err := o.QueryTable("employees").Filter("name", "ngoc").Limit(1).One(&dep)
	if err == nil {
		fmt.Println(dep.Name)
	}
	return

}
//GetAllEmployees retrieves all Employees matches certain condition. Returns empty list if
//no records exist
func GetAllEmployeesDefault(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(Employees))
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

	var l []Employees
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

func GetAllEmployees()(ml []Employees,err error){
	fmt.Println("model.getallemploess")

	o := orm.NewOrm()
	var maps []orm.Params
	//var lists []orm.ParamsList
	//var list orm.ParamsList
	//num, err := o.QueryTable("user").ValuesFlat(&list, "name")
	qs := o.QueryTable("employees")
	//cond := orm.NewCondition()
	//cond1:=cond.And("age__lt",30)
	num,err:=qs.Limit(3).OrderBy("id").Values(&maps)
	_=num
	if err == nil {
		fmt.Printf("Result Nums: %d\n", num)
		for _, m := range maps {
			fmt.Println(m["Id"],m["Name"])
		}
	}

	//if err == nil {
	//	fmt.Printf("Result Nums: %d\n", num)
	//	fmt.Printf("All User Names: %s", strings.Join(list, ", "))
	//}
	//if err==nil{
	//	for _,emp:=range l{
	//
	//		ml= append(ml, emp)
	//	}
	//}
	return ml,err
}
// UpdateEmployees updates Employees by Id and returns error if
// the record to be updated doesn't exist
func UpdateEmployeesById(m *Employees) (err error) {
	o := orm.NewOrm()
	v := Employees{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEmployees deletes Employees by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEmployees(id int) (err error) {
	o := orm.NewOrm()
	v := Employees{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Employees{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
