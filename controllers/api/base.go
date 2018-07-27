package api

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/conf"
	"strings"
	"strconv"
)

type ResponData struct {
	Status  bool 		`json:"status"`
	Msg  	string 		`json:"msg"`
	Data    interface{} `json:"data"`
}

type UserBase struct {
	beego.Controller
	Phone string
	IsReg bool
	Uid   int64
}

func (l *UserBase) Prepare()  {
	ret := l.Ctx.GetCookie(conf.CookieNmae)
	if len(ret) < 12{
		l.Data["json"] = ResponData{Status:false,Msg:"请先登录"}
		l.ServeJSON(true)
		return
	}
	info := strings.Split(ret,"|")
	if len(info) == 2{
		l.Phone = info[1]
		uid,err := strconv.ParseInt(info[0],10,64)
		if err != nil{
			l.Data["json"] = ResponData{Status:false,Msg:"Cookies不正确"}
			l.ServeJSON(true)
			return
		}
		l.Uid = uid
	}
}