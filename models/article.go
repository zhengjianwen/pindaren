package models

import (
	"time"
	"github.com/zhengjianwen/utils/log"
	"fmt"
)

// 评价表
type Article struct {
	Id 			int64 		`json:"id"`
	ClassifyId 	int64 		`json:"classify_id"` // 分类ID
	Title 		string 		`json:"title" xorm:"varchar(25)"` // 标题
	Phone       string 		`json:"-"`
	UserName    string 		`json:"username"`
	Content     string 		`json:"content" xorm:"varchar(10000)"` // 内容
	ImagesCount int64 		`json:"images_count" ` 	// 上传图片数
	ReadCount   int64 		`json:"read_count" ` 	// 阅读量
	Created 	time.Time 	`json:"created" xorm:"created"` // 创建时间
	Updated 	time.Time 	`json:"updated" xorm:"updated"` 		// 最后一次修改时间
}


func (c *Article) Validate() error {
	if c.Title == ""{
		return fmt.Errorf("标题不能为空")
	}
	if c.Phone == ""{
		return fmt.Errorf("未获取到用户，请先登录")
	}
	if len(c.Content) <= 6{
		return fmt.Errorf("内容过少")
	}
	c.ReadCount = 0
	u,err := new(UserInfo).GetForPhone(c.Phone)
	if err != nil{
		return fmt.Errorf("查询用户信息错误")
	}
	c.UserName = u.Name
	return nil
}

func (c *Article) Creat() error {
	_,err := Orm.InsertOne(c)
	if err != nil {

	}
	return nil
}

func (c *Article) List(p int) error {
	beans := make([]*Article,0)
	start := p * 10
	err := Orm.Desc("id").Limit(10,start).Find(beans)
	if err != nil{
		log.Debugf("[Article][list] %v",err)
		return fmt.Errorf("获取评价列表失败")
	}
	return nil
}
