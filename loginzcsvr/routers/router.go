// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import "github.com/astaxie/beego"
import "loginzcsvr/controllers"

func init() {
    beego.Router("/", &controllers.MainController{})

    ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/customer",
				beego.NSInclude(
					&controllers.CustomerController{},
				),
			),
		)
	beego.AddNamespace(ns)
}
