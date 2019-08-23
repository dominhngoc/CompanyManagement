// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"company_management/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/departments",
			beego.NSInclude(
				&controllers.DepartmentsController{},
			),
			beego.NSRouter("GetOne/:id",&controllers.DepartmentsController{},"get:GetOne"),
			beego.NSRouter("/getsAll",&controllers.DepartmentsController{},"get:GetAllDB"),
			beego.NSRouter("/Post",&controllers.DepartmentsController{},"post:Post"),
			beego.NSRouter("/edit/:id",&controllers.DepartmentsController{},"put:Put"),
			beego.NSRouter("/GetRelate",&controllers.DepartmentsController{},"get:GetRelateDepartment"),

		),

		beego.NSNamespace("/employees",
			beego.NSInclude(
				&controllers.EmployeesController{},

			),
			beego.NSRouter("/Delete/:id",&controllers.EmployeesController{},"get:Delete"),
			beego.NSRouter("/AddMany",&controllers.EmployeesController{},"get:AddMany"),
			beego.NSRouter("/AddOne",&controllers.EmployeesController{},"get:AddOne"),
			beego.NSRouter("/Post",&controllers.EmployeesController{},"post:Post"),
			beego.NSRouter("/GetAll",&controllers.EmployeesController{},"get:GetAll"),
			beego.NSRouter("/GetRelate",&controllers.EmployeesController{},"get:GetRelate"),

		),
	)
	beego.AddNamespace(ns)
}
