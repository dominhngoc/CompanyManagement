package controllers

import (
	"company_management/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

// EmployeesController operations for Employees
type EmployeesController struct {
	beego.Controller
}

// URLMapping ...
func (c *EmployeesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne/:id", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete/:id", c.Delete)
	c.Mapping("AddMany", c.AddMany)
	c.Mapping("AddOne", c.AddOne)
	c.Mapping("GetRelate",c.GetRelate)


}

// Post ...
// @Title Post
// @Description create Employees
// @Param	body		body 	models.Employees	true		"body for Employees content"
// @Success 201 {int} models.Employees
// @Failure 403 body is empty
// @router / [post]
func (c *EmployeesController) Post() {
	var v models.Employees
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddEmployees(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Employees by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Employees
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EmployeesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetEmployeesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Employees
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Employees
// @Failure 403
// @router / [get]
func (c *EmployeesController) GetRelate(){
	models.GetRelateModel()
}
func (c *EmployeesController) GetAll(){
	fmt.Println("getallemployees")
	l, err := models.GetAllEmployees()
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = l
		}
		c.ServeJSON()
}

func (c *EmployeesController) GetAlldefault() {

	println("test")

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	l, err := models.GetAllEmployeesDefault(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
// Put ...
// @Title Put
// @Description update the Employees
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Employees	true		"body for Employees content"
// @Success 200 {object} models.Employees
// @Failure 403 :id is not int
// @router /:id [put]

func (c *EmployeesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Employees{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateEmployeesById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
func (c *EmployeesController) AddOne() {
	fmt.Println("Addone runeed")
	v := models.Employees{
		Email: "ngoc41234",
		Age:   1,
		Name:  "ngoc",
		Id:    13,
	}
	models.AddOneEmployees(&v)
	//c.Data["json"]="name:ngoc"
	//c.ServeJSON()
	//}
}
func (c *EmployeesController) AddMany()  {
	fmt.Println("Controller runeed")
	users := []models.Employees{
		{Name: "slene",Id:31},
		{Name:"hieu",Id:32},
		{Name:"toan",Id:33},
	}
	models.AddManyEmployees(&users)
	//c.Data["json"]="name:ngoc"
	//c.ServeJSON()
}
// Delete ...
// @Title Delete
// @Description delete the Employees
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EmployeesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteEmployees(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
