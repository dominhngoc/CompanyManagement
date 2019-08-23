package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"],
        beego.ControllerComments{
            Method: "GetRelateDepartment",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["company_management/controllers:DepartmentsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["company_management/controllers:EmployeesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["company_management/controllers:EmployeesController"],
        beego.ControllerComments{
            Method: "GetRelate",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["company_management/controllers:EmployeesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["company_management/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["company_management/controllers:EmployeesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
