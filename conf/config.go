package conf

import "github.com/astaxie/beego"

var SessionName = "prd"
var CookieNmae = "pindaren"
var CookieAdminNmae = "prdadmin"

// 环境配置
func init()  {
	beego.BConfig.WebConfig.Session.SessionProvider="memory"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "./session"

}