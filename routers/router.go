package routers

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pdrgo/controllers/user"
	"github.com/zhengjianwen/pdrgo/controllers/utils"
	"github.com/zhengjianwen/pdrgo/controllers/pdr"
)

func init() {
    beego.Router("/", &pdr.Index{})
    beego.Router("/login.html", &pdr.Login{})
    beego.Router("/login/sms", &utils.Sms{})

    beego.Router("/user/index.html", &user.UserController{})
}
