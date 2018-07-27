package routers

import (
	"github.com/zhengjianwen/pindaren/controllers/admin"
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/controllers/api"
)

func init() {
	// admin
    beego.Router("/pindaren/login", &admin.LoginController{})
    beego.Router("/pindaren/logout", &admin.LogOutController{})
    beego.Router("/pindaren/index", &admin.IndexController{})
    beego.Router("/pindaren/users", &admin.UsersController{})

    //api
	beego.Router("/", &admin.LoginController{})
	beego.Router("/api/login", &api.Login{})
	beego.Router("/api/logout", &api.Logout{})

}
