package utils

import (
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pdrgo/controllers"
	"time"
	"crypto/rand"
	"github.com/zhengjianwen/pdrgo/http/cache"
	"fmt"
	"github.com/zhengjianwen/utils/log"
)

type Sms struct {
	beego.Controller
}

func (pdr *Sms)Post()  {
	phone := pdr.GetString("phone")
	data := controllers.Resp{Status:false}
	if phone == "" || len(phone) != 11{
		data.Msg = "手机号不正确"
	}else {
		code := GenToken()
		key := cache.KLoginRetry(phone)
		err := cache.Set(key,code)
		if err != nil{
			data.Msg = "获取失败"
		}else {
			//dx := utils.Sms{Phone:phone,Content:fmt.Sprintf("您的验证码是：%s。请不要把验证码泄露给其他人。",code)}
			//err := utils.SendSms(&dx)
			fmt.Println(phone + "短信验证码为："+ code)
			var err error
			if err != nil{
				data.Msg = err.Error()
			}else {
				data.Status = true
				data.Msg = "发送成功"
				log.Debugf("[sms] code:%s",code)
			}
		}
	}
	pdr.Data["json"] = data
	pdr.ServeJSON()
}

// genToken 生成随机字符串
func GenToken() string {
	const alphanum = "0123456789"
	var bytes = make([]byte, 6)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func DateStr() string {
	return time.Unix(time.Now().UTC().UnixNano() / 1000000000, 0).Format("20060102")
}