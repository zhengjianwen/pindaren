package admin

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/utils/log"
	"github.com/zhengjianwen/pindaren/models"
	"github.com/zhengjianwen/pindaren/conf"
	"github.com/zhengjianwen/pindaren/utils"
	"strings"
)

type AdminBase struct {
	beego.Controller
	IsLogin bool
	User models.Admin
}
type ResponData struct {
	Status  bool 		`json:"status"`
	Msg  	string 		`json:"msg"`
	Data    interface{} `json:"data"`
}

type ApiBase struct {
	beego.Controller
	methodGet bool          //是否GET请求
}

func (this *AdminBase) Prepare() {
	data := ResponData{Status:false,Msg:"请先登录"}
	str := this.Ctx.GetCookie(conf.CookieNmae)
	if str == ""{
		this.Data["json"] = data
		this.ServeJSONP()
		return
	}else {

		var admin models.Admin
		ret := utils.CookieDecode(str)
		data := strings.Split(ret,"|")
		if ret == "" || len(data) != 3{
			data := ResponData{Status:false,Msg:"请先登录"}
			this.Data["json"] = data

			this.ServeJSONP()
			return
		}

		this.User.Name = data[2]
		this.User.Phone = data[1]
		log.Debugf("[AdminBase] %s",admin)
		this.IsLogin = true
		this.Data["name"] = data[2]

	}
	log.Debugf("[Base][Prepare] %v",this)
}
