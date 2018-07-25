package models

import "time"

// 用户收藏表
type Collection struct {
	ID  		int64 `json:"id"`
	Aid 		int64 `json:"aid"` // 评价ID
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
}

