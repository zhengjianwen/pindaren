package models

import "time"

// 评价表
type Article struct {
	Id 			uint64 		`json:"id"`
	ClassifyId 	uint64 		`json:"classify_id"` // 分类ID
	Title 		string 		`json:"title" orm:"size(128)"`
	Content     string 		`json:"content" orm:"type(text)"` // 内容
	ImagesCount int64 		`json:"images_count" orm:"size(3)"` 	// 上传图片数
	ReadCount   int64 		`json:"read_count" orm:"size(15)"` 	// 阅读量
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
	Updated 	time.Time 	`json:"updated" orm:"auto_now;type(datetime)"` 		// 最后一次修改时间
}
