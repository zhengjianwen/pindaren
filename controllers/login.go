package controllers

type LoginController struct {
	Base
}

func (l *LoginController) Get() {
	l.TplName = "index.html"
}