package pdr

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pdrgo/http/cache"
	"github.com/zhengjianwen/pdrgo/models"
	"github.com/zhengjianwen/pindaren/conf"
)

type Login struct {
	beego.Controller
}

func (pdr *Login)Get()  {
	pdr.Data["xsrfdata"] = pdr.XSRFToken()
	pdr.TplName = "user/login.html"
}

func (pdr *Login)Post()  {
	phone := pdr.GetString("phone")
	code := pdr.GetString("code")
	key := cache.KLoginRetry(phone)
	val := cache.Cache.Get(key)
	if code == val{
		callback := pdr.GetString("callback")
		var user = models.UserInfo{Phone:phone}
		u,err := user.GetForPhone(phone)
		if err != nil{
			pdr.Data["err"] = "创建用户失败"
			pdr.Data["phone"] = phone
			pdr.TplName = "user/login.html"
			return
		}else {
			if u == nil{
				err = user.Save()
				if err != nil{
					pdr.Data["err"] = "用户注册失败"
					pdr.Data["phone"] = phone
					pdr.TplName = "user/login.html"
					return
				}
			}else {
				if u.IsLock{
					pdr.Data["err"] = "用户被锁定"
					pdr.Data["phone"] = phone
					pdr.TplName = "user/login.html"
					return
				}
			}
			pdr.Ctx.SetCookie(conf.CookieNmae,phone)

		}
		if callback != ""{
			pdr.Redirect(callback,301)
			return
		}else {
			pdr.Redirect("/",301)
		}
		return
	}
	pdr.Data["err"] = "验证码不正确"
	pdr.Data["phone"] = phone
	pdr.TplName = "user/login.html"





}