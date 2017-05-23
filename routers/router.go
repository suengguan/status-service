// @APIVersion 1.0.0
// @Title status-service API
// @Description status-service only serve status
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app-service/status-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/status",
			beego.NSInclude(
				&controllers.JobStatusController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
