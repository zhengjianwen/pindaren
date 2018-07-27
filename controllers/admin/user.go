package admin

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/models"
	"github.com/zhengjianwen/utils/log"
)

type UsersController struct {
	beego.Controller
}

func (l *UsersController) Get() {
	log.Debugf("123123123")
	user := new(models.Admin)
	user.Phone = "17710633221"
	new(models.Admin).Get("Phone")
	data := ResponData{Status:false}
	l.Data["json"] = data
	l.ServeJSON(true)
	return
}
