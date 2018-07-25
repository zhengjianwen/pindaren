package models

type SysInfo struct {
	Id   int64
	Name string   	// 网站名称
	Logo string   	// 图片地址

	Tel  string   	// 电话
	Ruser int64   	// 注册用户数
	SmsCount int64  // 发生短信数
}
