package models


type UserInfo struct {
	ID 			uint64 `json:"id" orm:"rel(fk);auto;unique;size(14)"`
	User 		*AuthUser `json:"-" orm:"reverse(one)"`

	Name 		string `json:"name" orm:"column(nickname);size(120);null;"` // 真实姓名
	IdCard 		string `json:"id_card" orm:"column(id_card);null;"` // 身份证号
	IsValidate	bool 	`json:"is_validate" orm:"null;default(0)"`
	Phone 		string `json:"phone" orm:"null;"` // 手机号
	Email 		string `json:"email" orm:"null;"` // 邮箱账号
	WxID  		string `json:"wx_id" orm:"null;"` // 微信号
	Qq    		string `json:"qq" orm:"null;"`    // QQ号
	Alipay 		string `json:"alipay" orm:"null;"` // 支付宝号

	School 		string `json:"school" orm:"null;"` // 学校名称
	City  		string `json:"city" orm:"null;"` // 城市
	Province 	string `json:"province" orm:"null;"` // 省份
}


