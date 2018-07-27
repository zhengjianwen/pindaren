package admin

import (
	"strings"
	"github.com/zhengjianwen/pindaren/models"
	"github.com/rongyungo/www/log"
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/utils"
	"fmt"
	"github.com/zhengjianwen/pindaren/conf"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	if l.GetSession("pdr") != nil{
		l.Redirect("/pindaren/index",301)
		return
	}
	l.TplName = "login/adminlogin.html"
}

func (l *LoginController) Post() {
	username := l.GetString("username")
	l.Data["username"] = username
	if username == ""{
		l.Data["msg"] = "用户名不能为空"
		l.TplName = "login/adminlogin.html"
		return
	}
	u := new(models.Admin)
	if strings.Contains(username, "@") {
		u.UserName = username
		u.Get("username")
	} else {
		u.Phone = username
		u.Get("phone")
	}
	if u.Id == 0{
		l.Data["msg"] = "管理员不存在"
		l.TplName = "login/adminlogin.html"
		return
	}
	log.Debugf("[LoginController] user:%v",u)

	pasw := l.GetString("password")
	pasw = models.EncryptPassword(pasw)
	if pasw == u.PassWord{

		token := utils.CookieEncode(fmt.Sprintf("%d|%s|%s",u.Id,u.Phone,u.Name))
		//l.SetSession("prd",u.Phone)
		log.Debugf("设置缓存：%s",token)
		l.Ctx.SetCookie(conf.CookieNmae,token)
		l.Redirect("/pindaren/index",301)
		return
	}else {
		l.Data["msg"] = "用户名或密码不正确"
	}
	l.TplName = "login/adminlogin.html"
}

