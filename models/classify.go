package models

import (
	"fmt"
	"github.com/zhengjianwen/utils/log"
)

// 分类信息表
type Classify struct {
	Id  		uint64 		`json:"id"`
	Name 		string 		`json:"name" xorm:"notnull"`
	Status      bool 		`json:"status" xorm:"default(1)"`
}
// 体验报告 购物晒单 求助 爆料

func (c *Classify) List() ([]*Classify,error) {
	beans := make([]*Classify,0)
	err := Orm.Where("status = ?",true).Find(&beans)
	if err != nil{
		return nil,fmt.Errorf("获取分类信息错误")
	}
	return beans,nil
}

func (c *Classify) Save() (error) {
	c.Status = true
	_,err := Orm.InsertOne(c)
	if err != nil{
		return fmt.Errorf("添加分类失败")
	}
	return nil
}

func (c *Classify) Del() (error) {
	c.Status = true
	_,err := Orm.Delete(c)
	if err != nil{
		log.Debugf("[Classify][del] %v",err)
		return fmt.Errorf("删除分类失败")
	}
	return nil
}

