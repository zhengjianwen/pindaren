package api

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/utils"
	"fmt"
	"github.com/zhengjianwen/pindaren/conf"
	"github.com/zhengjianwen/pindaren/models"
)

type Login struct {
	beego.Controller
}

func (l *Login)Get()  {
	// 获取短信
	phone := l.GetString("phone")
	if phone == "" || len(phone) != 11{
		l.Data["json"] = ResponData{Status:false,Msg:"手机号不能为空或手机号不正确"}
		l.ServeJSON()
		return
	}
	code := utils.GenCode()

	err := utils.SetCache(phone,code,60*3)
	if err != nil{
		l.Data["json"] = ResponData{Status:false,Msg:"发送验证码失败"}
		l.ServeJSON()
		return
	}
	content := fmt.Sprintf("您的验证码是：【%s】。请不要把验证码泄露给其他人。",code)
	utils.SendSms(&utils.Sms{Phone:phone,Content:content})
	l.Data["json"] = ResponData{Status:true,Msg:"发送验证码成功"}
	l.ServeJSON()
}

func (l *Login) Post() {
	data := ResponData{Status:false}
	l.Data["json"] = data

	phone := l.GetString("phone")
	key := l.GetString("code")
	if len(phone) != 11{
		data.Msg = "手机号不正确"
		l.Data["json"] = data
		l.ServeJSON()
		return
	}
	var  code  string
	err := utils.GetCache(phone,code)
	if err != nil{
		data.Msg = "验证码不存在"
		l.ServeJSON()
		return
	}
	if code != key{
		data.Msg = "验证码错误"
		l.ServeJSON()
		return
	}
	data.Status =true
	// 保存账号
	user := new(models.AuthUser)
	user.Phone = phone
	err = user.Save()
	if err !=nil{
		data.Msg = err.Error()
		l.ServeJSON()
		return
	}
	if user.Id < 0{
		data.Msg = "登录失败"
		l.ServeJSON()
		return
	}
	ret := fmt.Sprintf("%d|%s",user.Id,phone)
	l.Ctx.SetCookie(conf.CookieNmae,ret)
	data.Data = user
	l.ServeJSON()
}


type Logout struct {
	beego.Controller
}

func (l *Logout)Get()  {
	// 获取短信
	l.Ctx.SetCookie(conf.CookieNmae,"")
	l.Data["json"] = ResponData{Status:true,Msg:"退出成功"}
	l.ServeJSON()
}