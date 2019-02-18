// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"tianwei.pro/sam-web/controllers"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/portal",
			beego.NSInclude(
				&controllers.PortalController{},
			),
		),
		beego.NSNamespace("/system",
			beego.NSInclude(
				&controllers.SystemController{},
			)),
	)
	beego.AddNamespace(ns)
}
