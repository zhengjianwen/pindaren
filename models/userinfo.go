package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/zhengjianwen/utils/log"
	"fmt"
)

type UserInfo struct {
	Id 			uint64 `json:"id"`
	Name 		string `json:"name" orm:"column(nickname);size(120);null;"` // 真实姓名
	IdCard 		string `json:"id_card" orm:"column(id_card);null;"` // 身份证号
	IsValidate	bool 	`json:"is_validate" orm:"null;default(0)"`

	Email 		string `json:"email" orm:"null;"` // 邮箱账号
	WxID  		string `json:"wx_id" orm:"null;"` // 微信号
	Qq    		string `json:"qq" orm:"null;"`    // QQ号
	Alipay 		string `json:"alipay" orm:"null;"` // 支付宝号

	School 		string `json:"school" orm:"null;"` // 学校名称
	City  		string `json:"city" orm:"null;"` // 城市
	Province 	string `json:"province" orm:"null;"` // 省份
}

func (c *UserInfo) Validate() string {

	return ""
}

func (c *UserInfo) Creat() error {
	o := orm.NewOrm()
	o.Using("default")
	msg := c.Validate()
	if msg == ""{
		_,err := o.Insert(c)
		if err != nil{
			log.Errorf("[UserInfo][Creat] 创建失败 %v",err)
			return fmt.Errorf("创建失败")
		}
	}else {
		return fmt.Errorf(msg)
	}
	return nil
}

func (c *UserInfo) Update() error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Update(c)
	if err != nil{
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (c *UserInfo) Del() error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Update(c)
	if err != nil{
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (c *UserInfo) Get(id uint64) error {
	o := orm.NewOrm()
	o.Using("default")
	c.Id = id

	err := o.Read(&c)
	if err != orm.ErrNoRows{
		return fmt.Errorf("无此数据")
	}else if err == orm.ErrMissPK{
		return fmt.Errorf("无主键无法查询")
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
