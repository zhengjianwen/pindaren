package models

import "time"

type Images struct {
	Id      	int64     	`json:"id"`
	Aid     	int64     	`json:"aid"`                                      // 评价ID
	Position	int64 		`json:"position"` 								  // 第几张
	Path    	string 		`json:"path"`										// 图片路径
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
}
