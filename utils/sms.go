package utils

import (
	"strings"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"net/url"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
)

type Sms struct {
	Phone string
	Content string
}

// 您的验证码是：【变量】。请不要把验证码泄露给其他人。


// 支持多自定义版本的短信网关
func SendSms(sms *Sms) error {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := beego.AppConfig.DefaultString("smsuser","C64129556")	//查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := beego.AppConfig.DefaultString("smskey","124bc469e69a81f8e96b081371da6eaa") //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := sms.Phone
	_content := sms.Content
	_url := beego.AppConfig.DefaultString("smsurl","http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json")
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", _url, body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(string(data), err)
		return fmt.Errorf("短信发送失败")
	}
	return nil
}


func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}