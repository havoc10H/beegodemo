package routers

import (
	"quickstart/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	web.Router("/", &controllers.MainController{})
}
