package controllers

import "github.com/astaxie/beego"

type Base struct {
	beego.Controller
	methodGet bool          //是否GET请求
}

type ApiBase struct {
	beego.Controller
	methodGet bool          //是否GET请求
}

func (this *Base) Prepare() {
	//s := this.GetSession("pindaren")
	//if s == nil{
	//	this.StopRun()
	//}else {
	//
	//}

}

func (this *ApiBase) Prepare() {

}