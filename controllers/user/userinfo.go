package user

import "github.com/zhengjianwen/pdrgo/controllers"

type UserController struct {
	controllers.Base
}

func (pdr *UserController)Get()  {
	if pdr.IsLogin{
		pdr.Data["user"] = pdr.User
		pdr.TplName = "user/memberIndex.html"
	}else {
		pdr.Redirect("/login.html?callback=/user/index.html",301)
	}

}