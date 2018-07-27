package admin

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/conf"
)

type LogOutController struct {
	beego.Controller
}

func (l *LogOutController) Get() {
	l.Ctx.SetCookie(conf.CookieNmae,"")
	l.DelSession(conf.SessionName)
	l.Data["json"] = ResponData{Status:true,Msg:"退出成功"}
	l.ServeJSON(true)
	return
}