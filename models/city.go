package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/zhengjianwen/utils/log"
	"fmt"
)

type City struct {
	Id 			int64 	`json:"id"`
	Name 		string	`json:"name"` // 名称
	ParentId 	int64 	`json:"parent_id"` // 父级id
	Level 		int64 	`json:"level"` // 级别 0 省 1是市 2 县 3 乡
}

func (c *City) Validate() string {
	if c.Name == ""{
		return "名称不能为空"
	}
	if c.ParentId == 0{
		c.Level = 0
	}else {

	}

	return ""
}

func (c *City) Creat() error {
	o := orm.NewOrm()
	o.Using("default")
	msg := c.Validate()
	if msg == ""{
		_,err := o.Insert(c)
		if err != nil{
			log.Errorf("[City][Creat] 创建失败 %v",err)
			return fmt.Errorf("创建失败")
		}
	}else {
		return fmt.Errorf(msg)
	}
	return nil
}

func (c *City) Update() error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Update(c)
	if err != nil{
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (c *City) Del() error {
	o := orm.NewOrm()
	o.Using("default")
	_,err := o.Update(c)
	if err != nil{
		return fmt.Errorf("更新失败")
	}
	return nil
}

func (c *City) Get(id int64) error {
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

func (c *City) List(id,pid int64) ([]City,error) {
	o := orm.NewOrm()
	o.Using("default")
	sql := fmt.Sprintf("SELECT id, name,parent_id,level FROM city WHERE ")
	if pid != 0{
		sql += fmt.Sprintf("parent_id = %d ",pid)
	}
	if id != 0{
		sql += fmt.Sprintf("id = %d ",pid)
	}
	beans := make([]City,0)
	_,err := o.Raw(sql).QueryRows(&beans)
	if err != nil{
		return beans,fmt.Errorf("更新失败")
	}
	return beans,nil
}