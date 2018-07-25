package controllers

type LoginController struct {
	Base
}

func (l *LoginController) Get() {
	l.TplName = "index.html"
}

func (l *LoginController) Post() {
	phone := l.GetString("phone")
	//goback := l.GetString("goback")
	if len(phone) == 11{
		l.TplName = "user.html"
		l.SetSession("pindaren",phone)
		return
	}
	l.Data["msg"] = 1
	l.TplName = "index.html"
}

