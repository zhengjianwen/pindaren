package session

import "github.com/astaxie/beego/session"

var PdrSession *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"pingdaren",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	PdrSession, _ = session.NewManager("memory",sessionConfig)
	go PdrSession.GC()
}
