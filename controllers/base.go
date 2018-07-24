package controllers

import "github.com/astaxie/beego"

type Base struct {
	beego.Controller
	methodGet bool          //是否GET请求
}
