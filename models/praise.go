package models

import "time"

// 点赞信息表
type Praise struct {
	Id  		uint64 		`json:"id"`
	Aid 		uint64 		`json:"aid"` // 评价ID
	Cid 		uint64 		`json:"cid" orm:"default(0)"` // 回复ID
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
 }
