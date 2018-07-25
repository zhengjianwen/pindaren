package models

import "time"

// 回复表
type Comment struct {
	Id 			uint64 	`json:"id"`
	Aid 		uint64 	`json:"aid"`
	Uid 		uint64 	`json:"uid"`
	Content     string 	`json:"content" orm:"type(text)"` // 内容
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
	Updated 	time.Time 	`json:"updated" orm:"auto_now;type(datetime)"` 		// 最后一次修改时间
} 
