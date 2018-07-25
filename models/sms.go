package models

import "time"

type SendSms struct {
	Id 			int64 `json:"id"`
	Phone 		int64 `json:"phone"`
	Content 	string `json:"content"`
	Created 	time.Time 	`json:"created" orm:"auto_now_add;type(datetime)"` // 创建时间
}
