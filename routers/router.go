package routers

import (
	"github.com/astaxie/beego"
	"github.com/toophy/login/controllers"
)

func init() {
	beego.AutoRouter(&controllers.LocalController{})
}
