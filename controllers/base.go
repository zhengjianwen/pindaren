package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/zhengjianwen/pdrgo/conf"
	"fmt"
	"strings"
	"strconv"
	"github.com/zhengjianwen/pdrgo/models"
)

type Base struct {
	beego.Controller
	IsLogin bool
	UserID  int64
	Phone   string
	User    *models.UserInfo
}

func (pdr *Base)Prepare()  {
	pdr.Layout = "user/layout.html"
	pdr.LayoutSections = make(map[string]string)
	pdr.LayoutSections["HtmlHead"] = ""
	pdr.LayoutSections["Scripts"] = ""
	pdr.Data["xsrfdata"]=template.HTML(pdr.XSRFFormHTML())
	pdr.IsLogin = false
	session := pdr.GetSession(conf.SessionName)
	if session != nil{
		key := fmt.Sprintf("%v",session)
		list := strings.Split(key,"|")
		if len(list) == 2{
			uid,err :=  strconv.ParseInt(list[0],10,64)
			if err == nil{
				fmt.Println(uid)
				u,err := new(models.UserInfo).Get(uid,list[1])
				if err == nil && u != nil{
					pdr.User = u
					pdr.UserID = uid
					pdr.Phone = list[1]
					pdr.IsLogin = true
				}
			}
		}else {
			pdr.IsLogin = false
		}
	}
}