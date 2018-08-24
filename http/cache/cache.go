package cache

import (
	"github.com/astaxie/beego/cache"
	"time"
	"github.com/zhengjianwen/utils/log"
	"fmt"
)

var Cache cache.Cache
var MemoryCache cache.Cache

func init()  {
	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil{

	}
	Cache = bm
}

func Set(name,vaule string) error {
	return Cache.Put(name,vaule,time.Second * 60)
}

func Del(name string) error {
	return Cache.Delete(name)
}

func IsExist(name string) bool {
	return Cache.IsExist(name)
}

func Validate(name string,vaule interface{}) bool {
	ret := Cache.Get(name)
	if  ret == vaule{
		return true
	}
	log.Debugf("[Cache][Validate] RET:%v  vaule:%v",ret,vaule)
	return false
}

func KLoginRetry(phone string) string {
	now := time.Now().Unix()
	now = now - now%3600
	return fmt.Sprintf("pingdaren/user/%s/retry/%d", phone, now)
}