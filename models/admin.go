package models

import (
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/zhengjianwen/utils/log"
)

type Admin struct {
	Id  		uint64 		`json:"id" orm:"pk"`
	Name 		string 		`json:"nick_name" orm:"column(nickname);size(120);null"`
	UserName 	string 		`json:"username" orm:"column(username);index;size(120);null"` // 账号
	Phone 		string 		`json:"phone" orm:"index"` // 手机号
	PassWord 	string 		`json:"-" orm:"column(password);size(128);"`// 密码
	IsLock  	bool 		`json:"is_lock" orm:"default(1)"`    // 锁定

	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
	LastLogin  	time.Time 	`json:"lastlogin" orm:"type(datetime)"` // 最后一次登录时间
	Updated 	time.Time 	`json:"updated" orm:"auto_now;type(datetime)"` // 最后一次修改时间
}

func (u *Admin) Get(pk string) (error) {
	if u == nil{
		return fmt.Errorf("用户信息不能为空")
	}
	o := orm.NewOrm()
	o.Using("default")
	//
	err := o.Read(u,pk)
	if err != nil{
		log.Debugf("[Admin][Get] %v",err)
		return fmt.Errorf("获取用户信息失败")
	}

	return nil
}

func (u *Admin) Save() (error) {
	if u == nil{
		return fmt.Errorf("用户信息不能为空")
	}
	if msg := u.Validate();msg != ""{
		return fmt.Errorf(msg)
	}
	o := orm.NewOrm()
	o.Using("default")
	//
	_,err := o.InsertOrUpdate(u)
	if err != nil{
		log.Errorf("[Admin][Save] 数据错误")
		return fmt.Errorf("插入数据或更新失败")
	}
	return nil
}

func (u *Admin) Validate() (string) {
	if u == nil{
		return "用户信息不能为空"
	}
	if u.Phone == ""{
		return "用户手机号不能为空"
	}
	if u.PassWord == ""{
		return "用户密码不能为空"
	}
	if len(u.PassWord) < 32 || len(u.PassWord) > 128 {
		return "密码编码格式不正确"
	}
	return ""
}
