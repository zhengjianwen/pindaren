package models

import "time"

type AuthUser struct {
	ID  		uint64 		`json:"id" orm:"rel(fk);auto;unique;size(14)"`

	NickName 	string 		`json:"nick_name" orm:"column(nickname);size(120)"`
	UserName 	string 		`json:"username" orm:"column(username);unique;size(120);"` // 账号
	PassWord 	string 		`json:"-" orm:"column(password);size(128);"`// 密码
	IsLock  	bool 		`json:"is_lock"`    // 是否锁定
	UserInfo    *UserInfo   `orm:"rel(one)"`

	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
	LastLogin  	time.Time 	`json:"lastlogin" orm:"type(datetime)"` // 最后一次登录时间
	Updated 	time.Time 	`json:"updated" orm:"auto_now;type(datetime)"` // 最后一次修改时间
}
func (u *AuthUser) TableIndex() [][]string {
	return [][]string{
		[]string{"username","nickname"},
	}
}

