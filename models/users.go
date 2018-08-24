package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/zhengjianwen/utils/log"
	"fmt"
	"time"
)

type UserInfo struct {
	//Id  		int64 		`json:"id" orm:"pk"`
	Phone 		string 		`json:"phone" xorm:"index"` // 手机号
	Name 		string 		`json:"nickname" orm:"column(nickname);size(120);null"` // 别名

	IsLock  	bool 		`json:"is_lock" orm:"default(1)"`    // 锁定

	Created 	time.Time 	`json:"created" xorm:"created created"` // 创建时间
	LastLogin  	time.Time 	`json:"lastlogin" xorm:""` // 最后一次登录时间
	Updated 	time.Time 	`json:"updated" xorm:"updated updated"` // 最后一次修改时间

	CnName 		string 		`json:"cnname" orm:"column(cnname);size(120);null;"` // 真实姓名
	IdCard 		string 		`json:"-" orm:"column(id_card);null;"` // 身份证号
	IsValidate	bool 		`json:"is_validate" orm:"null;default(0)"`   // 是否认证

	Email 		string 		`json:"email" orm:"null;"` // 邮箱账号
	WxID  		string 		`json:"wx_id" orm:"null;"` // 微信号
	Qq    		string 		`json:"qq" orm:"null;"`    // QQ号
	Alipay 		string 		`json:"alipay" orm:"null;"` // 支付宝号

	School 		string 		`json:"school" orm:"null;"` // 学校名称
	City  		string 		`json:"city" orm:"null;"` // 城市
	Province 	string 		`json:"province" orm:"null;"` // 省份
}

func (c *UserInfo) Validate() error {
	if len(c.Phone) != 11{
		return fmt.Errorf("手机号不正确")
	}

	return nil
}

func (c *UserInfo) Save() error {
	msg := c.Validate()
	if msg == nil{
		_,err := Orm.InsertOne(c)
		if err != nil{
			log.Errorf("[UserInfo][Creat] 创建失败 %v",err)
			return fmt.Errorf("创建失败")
		}
	}else {
		return msg
	}
	return nil
}

func (c *UserInfo) Update() error {
	if c.Phone == "" || len(c.Phone) != 11{
		return fmt.Errorf("验证信息失败")
	}
	_,err := Orm.Where("phone = ?",c.Phone).Update(c)
	if err != nil{
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (c *UserInfo) Get(id int64,phone string) (*UserInfo,error) {
	has,err := Orm.Where("id = ? and phone = ?",id,phone).Get(c)
	if err != nil{
		return nil,fmt.Errorf("查询用户失败")
	}
	if !has{
		return nil,nil
	}
	return c,nil
}

func (c *UserInfo) GetForPhone(phone string) (*UserInfo,error) {
	has,err := Orm.Where("phone = ?",phone).Get(c)
	if err != nil{
		return nil,fmt.Errorf("查询用户失败")
	}
	if !has{
		return nil,nil
	}
	return c,nil
}

func (c *UserInfo) PhoneIsExect(phone string) error {
	n,err := Orm.Where("phone = ?").Count(c)
	if err != nil{
		return fmt.Errorf("查询手机号是否存在错误")
	}
	if n > 0{
		return fmt.Errorf("手机号已存在")
	}
	return nil
}

func (c *UserInfo) List() ([]UserInfo,error) {
	o := orm.NewOrm()
	o.Using("default")
	sql := fmt.Sprintf("SELECT * FROM UserInfo;")
	beans := make([]UserInfo,0)
	_,err := o.Raw(sql).QueryRows(&beans)
	if err != nil{
		return beans,fmt.Errorf("更新失败")
	}
	return beans,nil
}
