package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-service/status-service/controllers:JobStatusController"] = append(beego.GlobalControllerRouter["app-service/status-service/controllers:JobStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/job/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
