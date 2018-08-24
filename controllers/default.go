package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Resp struct {
	Status bool `json:"status"`
	Msg    string `json:"msg"`
	Code   uint64 `json:"code"`  
	Data   interface{} `json:"data"`
} 