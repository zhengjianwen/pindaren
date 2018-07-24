package routers

import (
	"github.com/zhengjianwen/pindaren/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
